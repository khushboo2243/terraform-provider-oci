---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_virtual_circuit"
sidebar_current: "docs-oci-datasource-core-virtual_circuit"
description: |-
  Provides details about a specific Virtual Circuit in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_virtual_circuit
This data source provides details about a specific Virtual Circuit resource in Oracle Cloud Infrastructure Core service.

Gets the specified virtual circuit's information.

## Example Usage

```hcl
data "oci_core_virtual_circuit" "test_virtual_circuit" {
	#Required
	virtual_circuit_id = "${oci_core_virtual_circuit.test_virtual_circuit.id}"
}
```

## Argument Reference

The following arguments are supported:

* `virtual_circuit_id` - (Required) The OCID of the virtual circuit.


## Attributes Reference

The following attributes are exported:

* `bandwidth_shape_name` - The provisioned data rate of the connection.
* `bgp_management` - BGP management option. 
* `bgp_session_state` - The state of the BGP session associated with the virtual circuit.
* `compartment_id` - The OCID of the compartment containing the virtual circuit.
* `cross_connect_mappings` - An array of mappings, each containing properties for a cross-connect or cross-connect group that is associated with this virtual circuit. 
	* `bgp_md5auth_key` - The key for BGP MD5 authentication. Only applicable if your system requires MD5 authentication. If empty or not set (null), that means you don't use BGP MD5 authentication. 
	* `cross_connect_or_cross_connect_group_id` - The OCID of the cross-connect or cross-connect group for this mapping. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider). 
	* `customer_bgp_peering_ip` - The BGP IP address for the router on the other end of the BGP session from Oracle. Specified by the owner of that router. If the session goes from Oracle to a customer, this is the BGP IP address of the customer's edge router. If the session goes from Oracle to a provider, this is the BGP IP address of the provider's edge router. Must use a /30 or /31 subnet mask.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IP addresses.

		Example: `10.0.0.18/31` 
	* `oracle_bgp_peering_ip` - The IP address for Oracle's end of the BGP session. Must use a /30 or /31 subnet mask. If the session goes from Oracle to a customer's edge router, the customer specifies this information. If the session goes from Oracle to a provider's edge router, the provider specifies this.

		There's one exception: for a public virtual circuit, Oracle specifies the BGP IP addresses.

		Example: `10.0.0.19/31` 
	* `vlan` - The number of the specific VLAN (on the cross-connect or cross-connect group) that is assigned to this virtual circuit. Specified by the owner of the cross-connect or cross-connect group (the customer if the customer is colocated with Oracle, or the provider if the customer is connecting via provider).  Example: `200` 
* `customer_bgp_asn` - The BGP ASN of the network at the other end of the BGP session from Oracle. If the session is between the customer's edge router and Oracle, the value is the customer's ASN. If the BGP session is between the provider's edge router and Oracle, the value is the provider's ASN. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `gateway_id` - The OCID of the customer's [Dynamic Routing Gateway (DRG)](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Drg) that this virtual circuit uses. Applicable only to private virtual circuits. 
* `id` - The virtual circuit's Oracle ID (OCID).
* `oracle_bgp_asn` - The Oracle BGP ASN.
* `provider_service_id` - The OCID of the service offered by the provider (if the customer is connecting via a provider). 
* `provider_state` - The provider's state in relation to this virtual circuit (if the customer is connecting via a provider). ACTIVE means the provider has provisioned the virtual circuit from their end. INACTIVE means the provider has not yet provisioned the virtual circuit, or has de-provisioned it. 
* `public_prefixes` - For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to advertise across the connection. Each prefix must be /24 or less specific. 
* `reference_comment` - Provider-supplied reference information about this virtual circuit (if the customer is connecting via a provider). 
* `region` - The Oracle Cloud Infrastructure region where this virtual circuit is located. 
* `service_type` - Provider service type. 
* `state` - The virtual circuit's current state. For information about the different states, see [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm). 
* `time_created` - The date and time the virtual circuit was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `type` - Whether the virtual circuit supports private or public peering. For more information, see [FastConnect Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/fastconnect.htm). 

