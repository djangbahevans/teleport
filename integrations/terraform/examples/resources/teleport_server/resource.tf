resource "teleport_server" "ssh_agentless" {
  version  = "v2"
  sub_kind = "openssh"
  // Name is not required for servers, this is a special case.
  // When a name is not set, an UUID will be generated by Teleport and
  // imported back into Terraform.
  // Giving unique IDs to servers allows UUID-based dialing (as opposed to
  // host-based dialing and IP-based dialing) which is more robust than its
  // counterparts as it can point to a specific server if multiple servers
  // share the same hostname/ip.
  spec = {
    addr     = "127.0.0.1:22"
    hostname = "test.local"
  }
}

resource "teleport_server" "ssh_agentless_eice" {
  version  = "v2"
  sub_kind = "openssh-ec2-ice"
  metadata = {
    // It is recommended to put the account and instance ID as a name for EC2 Instance Connect
    // When dialing to this instance, teleport will detect that this is an
    // AWS instance ID an will contact this specific instance. This is more
    // robust than host-based and IP-based dialing (because several server
    // can have similar hostnames).
    name = "123456789012-i-0123456789abcdef"
  }
  spec = {
    addr     = "127.0.0.1:22"
    hostname = "test.local"

    cloud_metadata = {
      aws = {
        account_id  = "123"
        instance_id = "123"
        region      = "us-east-1"
        vpc_id      = "123"
        integration = "foo"
        subnet_id   = "123"
      }
    }
  }
}