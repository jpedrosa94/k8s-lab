output "eks_cluster_name" {
  value = module.eks.cluster_name
}

output "eks_cluster_endpoint" {
  value = module.eks.cluster_endpoint
}

output "karpenter_queue_name" {
  value = module.karpenter.queue_name
}

output "karpenter_node_iam_role" {
  value = module.karpenter.node_iam_role_name
}

output "karpenter_instance_profile" {
  value = module.karpenter.instance_profile_name
}