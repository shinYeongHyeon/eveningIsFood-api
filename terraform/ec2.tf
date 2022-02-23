resource "aws_instance" "eif-api-dev" {
  ami = "ami-08508144e576d5b64"
  instance_type = "t2.micro"

  tags = {
    Name = "eif-api-dev"
  }
}