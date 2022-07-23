# Fork of Terraform Provider for AWS

Implementing AWS terraform datasources that are currently not available in official AWS terraform Provider

## Available DataSources

### RDS

- aws_db_instances (filter RDS instances based on provided resource tags)
- aws_rds_clusters (filter RDS clusters based on provided resource tags)

## Usage

```hcl
terraform {
  required_providers {
    awscust = {
      version = "5.1.2"
      source  = "msalman899/aws"
    }
  }
}

provider "awscust" {
    region = "eu-west-1"
    profile = "my-account"
}

#---------------------------
# aws_db_instances
#---------------------------

# Below data source will return all RDS instances that satisfy given tag's key-value criteria

data "aws_db_instances" "database" {
  provider = awscust
  filter {
    name = "team"
    values = ["value1","value2","value3"]
  }

  filter {
    name = "tribe"
    values = ["value1","value2"]
  }
  
  filter {
    name = "tagkey"
    values = ["tagvalue"]
  }
}

#---------------------------
# aws_rds_clusters
#---------------------------

# Below data source will return all RDS clusters that satisfy given tag's key-value criteria

data "aws_rds_clusters" "database" {
  provider = awscust
  filter {
    name = "tagkey"
    values = ["tagvalue"]
  }
}
```

# Terraform Provider for AWS

[![Forums][discuss-badge]][discuss]

[discuss-badge]: https://img.shields.io/badge/discuss-terraform--aws-623CE4.svg?style=flat
[discuss]: https://discuss.hashicorp.com/c/terraform-providers/tf-aws/

- Website: [terraform.io](https://terraform.io)
- Tutorials: [learn.hashicorp.com](https://learn.hashicorp.com/terraform?track=getting-started#getting-started)
- Forum: [discuss.hashicorp.com](https://discuss.hashicorp.com/c/terraform-providers/tf-aws/)
- Chat: [gitter](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing List: [Google Groups](http://groups.google.com/group/terraform-tool)

The Terraform AWS provider is a plugin for Terraform that allows for the full lifecycle management of AWS resources.
This provider is maintained internally by the HashiCorp AWS Provider team.

Please note: We take Terraform's security and our users' trust very seriously. If you believe you have found a security issue in the Terraform AWS Provider, please responsibly disclose by contacting us at security@hashicorp.com.

## Quick Starts

- [Using the provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)
- [Provider development](docs/contributing)

## Documentation

Full, comprehensive documentation is available on the Terraform website:

https://terraform.io/docs/providers/aws/index.html

## Roadmap

Our roadmap for expanding support in Terraform for AWS resources can be found in our [Roadmap](ROADMAP.md) which is published quarterly.

## Frequently Asked Questions

Responses to our most frequently asked questions can be found in our [FAQ](docs/contributing/faq.md )

## Contributing

The Terraform AWS Provider is the work of thousands of contributors. We appreciate your help!

To contribute, please read the contribution guidelines: [Contributing to Terraform - AWS Provider](docs/contributing)
