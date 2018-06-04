# Create your IAM Role

```
aws iam create-role --role-name ec2-role --assume-role-policy-document file://instance-role-trust-policy.json --profile private
aws iam attach-role-policy --policy-arn arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role --role-name ec2-role --profile private
aws iam attach-role-policy --policy-arn arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole --role-name ec2-role --profile private
aws iam create-instance-profile --instance-profile-name ec2-profile --profile private
aws iam add-role-to-instance-profile --role-name ec2-role --instance-profile-name ec2-profile --profile private
```

# Generate your ECS KeyPair

```
aws ec2 create-key-pair --key-name ec2-key --query 'KeyMaterial' --output text --profile private > ec2-key.pem
```

# make a file with the output

```
chmod 600 ec2-key.pem
ssh-add ec2-key.pem
```

# Create your ECS Security Group

```
aws ec2 create-security-group --group-name ec2-security-group --description "EC2 security group" --profile private
aws ec2 authorize-security-group-ingress --group-id sg-[YOUR GROUP ID] --protocol tcp --port 22 --cidr 0.0.0.0/0 --profile private
aws ec2 authorize-security-group-ingress --group-id sg-[YOUR GROUP ID] --protocol tcp --port 80 --cidr 0.0.0.0/0 --profile private
```

# Generate Your Cluster

```
aws ecs create-cluster --cluster-name taquito --profile private
```

# Set the Default VPC (make sure you have a default VPC)

```
aws ec2 describe-vpcs --profile private
```

# Spin up your EC2 Instance

```
aws ssm get-parameters --names /aws/service/ecs/optimized-ami/amazon-linux/recommended --profile private
aws ec2 run-instances --iam-instance-profile Name=ec2-profile --image-id ami-d2f489aa --count 1 --instance-type t2.micro --key-name ec2-key --security-group-ids sg-[YOUR GROUP ID] --user-data file://user-data.txt --associate-public-ip-address --profile private
```

Check here to see if the image-id is an image in your region: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ECS_AWSCLI_EC2.html#AWSCLI_EC2_launch_container_instance.

# Check Your EC2 Instance Booted

```
aws ec2 describe-instances --instance-ids i-[YOUR INSTANCE ID] --profile private
ssh -i ec2-key ec2-user@ec2-34-219-49-94.us-west-2.compute.amazonaws.com
```

# Once the instance boots and the user-data script runs it should attach to the cluster

```
aws ecs list-container-instances --cluster taquito --profile private
```
