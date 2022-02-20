resource "aws_db_instance" "default" {
  allocated_storage    = 10
  engine               = "aurora-mysql"
  engine_version       = "8.0"
  instance_class       = "db.t3.micro"
  db_name              = "eif"
  username             = var.db_username
  password             = var.db_password
  parameter_group_name = "default.mysql8.0"
  apply_immediately    = true
  character_set_name   = "utf8mb4"
  timezone             = "KST"
}
