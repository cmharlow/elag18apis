1. Create an IAM Role for the instance
```
aws --profile profile-name iam create-role --role-name ec2-role --assume-role-policy-document file://instance-role-trust-policy.json
```

2. Attach needed policies to the role
```
aws --profile profile-name iam attach-role-policy --role-name ec2-role --policy-arn arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role
aws --profile profile-name iam attach-role-policy --role-name ec2-role --policy-arn arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole
```

3. Create an Instance Profile & attach the IAM Role to it
```
aws --profile profile-name iam create-instance-profile --instance-profile-name ec2-profile
aws --profile profile-name iam add-role-to-instance-profile --role-name ec2-role --instance-profile-name ec2-profile
```

4. Create a SSH KeyPair for the instance and make a file with the KeyMaterial output. The file also needs to be Read by the owner only and added to the local ssh client.
```
aws --profile profile-name ec2 create-key-pair --key-name ec2-key
# make a file with the output
chmod 600 ec2-key
ssh-add ec2-key
```

5. Make sure you have a default VPC and make note of its ID
```
aws --profile profile-name ec2 describe-vpcs
```

6. Create a Security Group for the instance in the default VPC and make note of its ID
```
aws --profile profile-name ec2 create-security-group --group-name ec2-security-group --description "EC2 security group" --vpc-id ${vpc-id}
```

7. Add ingress rules to the Security Group
```
aws --profile profile-name ec2 authorize-security-group-ingress --group-id ${security-group-id} --protocol tcp --port 22 --cidr 0.0.0.0/0
aws --profile profile-name ec2 authorize-security-group-ingress --group-id ${security-group-id} --protocol tcp --port 80 --cidr 0.0.0.0/0
```

8. Create the ECS Cluster
```
aws --profile profile-name ecs create-cluster --cluster-name taquito
```

9. Query the System Manager API for the ECS optimized AMI and make a note of the AMI ID
```
aws --profile profile-name ssm get-parameters --names /aws/service/ecs/optimized-ami/amazon-linux/recommended
```

10. Spin up an EC2 Instance and make note of the Instance ID
```
aws --profile profile-name ec2 run-instances --iam-instance-profile Name=ec2-profile --image-id ${ami-id} --count 1 --instance-type t2.micro --key-name ec2-key --security-group-ids ${security-group-i} --user-data file://user-data.txt --associate-public-ip-address
```

11. After the instance has initialized, retrive its public DNS
```
aws --profile profile-name ec2 describe-instances --instance-ids ${instance-id}
```

12. SSH onto the Instance and check out the logs files in `/var/log/ecs`
```
ssh -i ec2-key ec2-user@ec2-${public-ip}.us-west-2.compute.amazonaws.com
```

13. Confirm the Instance attached to the cluster
```
aws --profile profile-name ecs list-container-instances --cluster taquito
```

14. Register the TAQUITO task definition and make note of taskDefinitionArn
```
aws --profile profile-name ecs register-task-definition --cli-input-json file://task-definition.json
```

15. Create the TAQUITO Service
```
aws --profile profile-name ecs create-service --service-name taquito --cluster taquito --task-definition ${task-definition-arn} --desired-count 1
```
