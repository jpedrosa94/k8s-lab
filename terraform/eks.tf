locals {
  name   = "ex-${basename(path.cwd)}"
  region = "us-east-1"

  vpc_cidr = "10.1.0.0/16"
  azs      = slice(data.aws_availability_zones.available.names, 0, 3)
}

module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "~> 20.11"

  cluster_name                             = local.name
  cluster_version                          = "1.30"
  enable_cluster_creator_admin_permissions = true
  cluster_endpoint_public_access           = true

  cluster_addons = {
    coredns                = {}
    eks-pod-identity-agent = {}
    kube-proxy             = {}
    vpc-cni                = {}
  }

  vpc_id                   = module.vpc.vpc_id
  subnet_ids               = module.vpc.private_subnets
  control_plane_subnet_ids = module.vpc.intra_subnets

  eks_managed_node_groups = {
    karpenter = {
      ami_type       = "AL2023_x86_64_STANDARD"
      instance_types = ["m5.large"]

      min_size     = 3
      max_size     = 5
      desired_size = 3

      #       taints = {
      #         # This Taint aims to keep just EKS Addons and Karpenter running on this MNG
      #         # The pods that do not tolerate this taint should run on nodes created by Karpenter
      #         addons = {
      #           key    = "CriticalAddonsOnly"
      #           value  = "true"
      #           effect = "NO_SCHEDULE"
      #         }
      #       }
    }
  }
  cluster_security_group_additional_rules = {
    ingress_self_all = {
      description = "Node to node all ports/protocols"
      protocol    = "-1"
      from_port   = 0
      to_port     = 0
      type        = "ingress"
      self        = true
    }
  }
}
