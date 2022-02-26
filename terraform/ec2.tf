resource "aws_instance" "eif_api_dev" {
  ami                         = "ami-08508144e576d5b64"
  instance_type               = "t2.micro"
  vpc_security_group_ids      = [aws_security_group.dev_web.id]
  key_name                    = aws_key_pair.eif_web.key_name
  subnet_id                   = module.vpc.public_subnets[0]

  tags = {
    Name = "eif-api-dev"
  }
}

resource "aws_eip" "eif_api_dev" {
  vpc                       = true
  instance                  = aws_instance.eif_api_dev.id
  associate_with_private_ip = aws_instance.eif_api_dev.private_ip
}
