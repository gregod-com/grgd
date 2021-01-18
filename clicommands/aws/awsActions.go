package aws

import (
	"io/ioutil"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/gregod-com/grgd/controller/helper"
	"github.com/gregod-com/grgd/interfaces"
	"github.com/urfave/cli/v2"
)

// ADBConnector ...
func ADBConnector(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	logger := core.GetLogger()

	awsregion, ok := c.App.Metadata["AWS-REGION"].(string)
	if !ok {
		logger.Fatal("AWS-REGION not defined")
	}
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsregion),
		Credentials: credentials.NewSharedCredentials("", "default"),
	})
	if err != nil {
		logger.Fatal(err.Error())
	}

	svc := rds.New(sess)
	result, err := svc.DescribeDBInstances(nil)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rds.ErrCodeDBInstanceNotFoundFault:
				logger.Fatalf(rds.ErrCodeDBInstanceNotFoundFault, aerr.Error())
			default:
				logger.Fatalf(aerr.Error())
			}
		} else {
			logger.Fatalf(err.Error())
		}
		return err
	}

	UI.Println("The database is currently available at:\n")
	UI.Println(*result.DBInstances[0].Endpoint.Address)
	UI.Println("\nExecuting nodeconnector command...\n")
	return ANodeConnector(c)
}

// ANodeConnector ...
func ANodeConnector(c *cli.Context) error {
	core := helper.GetExtractor().GetCore(c)
	UI := core.GetUI()
	logger := core.GetLogger()
	var fsm interfaces.IFileSystemManipulator
	core.Get(&fsm)

	heads := []string{"Nr", "Instance ID", "IP", "Avail. Zone", "DNS Name"}
	rows := [][]string{}

	awsregion, ok := c.App.Metadata["AWS-REGION"].(string)
	if !ok {
		logger.Fatal("AWS-REGION not defined")
	}
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsregion),
		Credentials: credentials.NewSharedCredentials("", "default"),
	})
	if err != nil {
		logger.Fatal(err.Error())
	}

	ec2Svc := ec2.New(sess)
	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		logger.Fatal("Error", err)
	}
	for k, res := range result.Reservations {
		instanceMetadata := []string{}
		for _, instance := range res.Instances {
			instanceMetadata = append(instanceMetadata, strconv.Itoa(k+1))
			instanceMetadata = append(instanceMetadata, *instance.InstanceId)
			instanceMetadata = append(instanceMetadata, *instance.PublicIpAddress)
			instanceMetadata = append(instanceMetadata, *instance.Placement.AvailabilityZone)
			instanceMetadata = append(instanceMetadata, *instance.PublicDnsName)
		}
		rows = append(rows, instanceMetadata)
	}

	UI.PrintTable(heads, rows, c)
	UI.Println()
	if nrArgs := c.Args().Len(); nrArgs != 1 {
		logger.Warn("Please enter the node number as an integer (i.e. '1', '2', '9')")
		return cli.NewExitError("Wrong number of arguments", 2)
	}
	if nrArgs := c.Args().Len(); nrArgs == 1 {
		instanceNr, err := strconv.Atoi(c.Args().First())
		if err != nil {
			logger.Warn("Please enter the node number as an integer (i.e. '1', '2', '9')")
			return cli.NewExitError("Wrong argument format", 3)
		}

		if instanceNr > len(rows) {
			UI.Println("\n\n\nThere are only " + strconv.Itoa(len(rows)) + " nodes available.")
			logger.Fatal("Please enter the node number as an integer (i.e. '1', '2', '9')")
		}
		// reduce by 1 to match slice index (UI prints nodes as 1,2,3... (not 0,1,2...))
		instanceNr--

		az := rows[instanceNr][3]
		instanceID := rows[instanceNr][1]
		IP := rows[instanceNr][2]
		pubkey, err := ioutil.ReadFile(fsm.HomeDir("/.ssh/id_rsa.pub"))
		if err != nil {
			logger.Fatal(err.Error())
		}
		logger.Debugf("Found instance %s\n", rows[instanceNr][1])

		svc := ec2instanceconnect.New(sess)
		input := &ec2instanceconnect.SendSSHPublicKeyInput{
			AvailabilityZone: aws.String(az),
			InstanceId:       aws.String(instanceID),
			InstanceOSUser:   aws.String("ec2-user"),
			SSHPublicKey:     aws.String(string(pubkey)),
		}

		result, err := svc.SendSSHPublicKey(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case ec2instanceconnect.ErrCodeAuthException:
					logger.Fatal(ec2instanceconnect.ErrCodeAuthException, aerr.Error())
				case ec2instanceconnect.ErrCodeInvalidArgsException:
					logger.Fatal(ec2instanceconnect.ErrCodeInvalidArgsException, aerr.Error())
				case ec2instanceconnect.ErrCodeServiceException:
					logger.Fatal(ec2instanceconnect.ErrCodeServiceException, aerr.Error())
				case ec2instanceconnect.ErrCodeThrottlingException:
					logger.Fatal(ec2instanceconnect.ErrCodeThrottlingException, aerr.Error())
				case ec2instanceconnect.ErrCodeEC2InstanceNotFoundException:
					logger.Fatal(ec2instanceconnect.ErrCodeEC2InstanceNotFoundException, aerr.Error())
				default:
					logger.Fatal(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				logger.Fatal(err.Error())
			}
		}

		logger.Debug(result)
		if *result.Success {
			UI.Printf("You can now connect to the ec2 instance at %s with user 'ec2-user'. Like this:\n", az)
			UI.Println("\nssh ec2-user@" + IP + "\n")
			UI.Println("!!! ssh key will be invalidated after 60 seconds !!!")
			UI.Println("!!! having a timeout? => your public IP might not be set correctly in the security group !!!")
			UI.Println("!!! make changes to the security group via AWS WebUI !!!\n")
		}
	}
	return nil
}
