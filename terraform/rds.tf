resource "aws_db_instance" "eif_dev" {
  identifier                = "eif-rds-dev"
  allocated_storage         = 100
  engine                    = "mysql"
  engine_version            = "8.0.28"
  instance_class            = "db.t3.medium"
  db_name                   = "eif"
  username                  = var.db_username
  password                  = var.db_password
  parameter_group_name      = "default.mysql8.0"
  apply_immediately         = true
  publicly_accessible       = true
  final_snapshot_identifier = "eif-rds-dev-final-snapshot"
  db_subnet_group_name      = module.vpc.database_subnet_group_name
  vpc_security_group_ids    = [aws_security_group.dev_rds.id]
}
