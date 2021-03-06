// This file was auto-generated from a YAML file

package v1_0

func init() {
	Profile["/tosca/simple-for-nfv/1.0/data.yaml"] = `
# Modified from a file that was distributed with this NOTICE:
#
#   Apache AriaTosca
#   Copyright 2016-2017 The Apache Software Foundation
#
#   This product includes software developed at
#   The Apache Software Foundation (http://www.apache.org/).

tosca_definitions_version: tosca_simple_yaml_1_2

data_types:

  tosca.datatypes.nfv.L2AddressData:
    # ERRATUM: TBD
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0-csd04]'
      citation_location: 5.3.1

  tosca.datatypes.nfv.L3AddressData:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0-csd04]'
      citation_location: 5.3.2
    description: >-
      The L3AddressData type is a complex TOSCA data type used to describe L3AddressData information
      element as defined in [ETSI GS NFV-IFA 011], it provides the information on the IP addresses
      to be assigned to the connection point instantiated from the parent Connection Point
      Descriptor.
    derived_from: tosca.datatypes.Root
    properties:
      ip_address_assignment:
        description: >-
          Specify if the address assignment is the responsibility of management and orchestration
          function or not. If it is set to True, it is the management and orchestration function
          responsibility.
        type: boolean
        required: true
      floating_ip_activated:
        description: Specify if the floating IP scheme is activated on the Connection Point or not.
        type: boolean
        required: true
      ip_address_type:
        description: >-
          Define address type. The address type should be aligned with the address type supported by
          the layer_protocol properties of the parent VnfExtCpd.
        type: string
        required: false
        constraints:
        - valid_values: [ ipv4, ipv6 ]
      number_of_ip_address:
        description: >-
          Minimum number of IP addresses to be assigned.
        type: integer
        required: false

  tosca.datatypes.nfv.AddressData:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0-csd04]'
      citation_location: 5.3.3
    description: >-
      The AddressData type is a complex TOSCA data type used to describe AddressData information
      element as defined in [ETSI GS NFV-IFA 011], it provides information on the addresses to be
      assigned to the connection point(s) instantiated from a Connection Point Descriptor.
    derived_from: tosca.datatypes.Root
    properties:
      address_type:
        description: >-
          Describes the type of the address to be assigned to the connection point instantiated from
          the parent Connection Point Descriptor. The content type shall be aligned with the address
          type supported by the layerProtocol property of the parent Connection Point Descriptor.
        type: string
        required: true
        constraints:
        - valid_values: [ mac_address, ip_address ]
      l2_address_data:
        # Shall be present when the addressType is mac_address.
        description: >-
          Provides the information on the MAC addresses to be assigned to the connection point(s)
          instantiated from the parent Connection Point Descriptor.
        type: tosca.datatypes.nfv.L2AddressData # Empty in "GS NFV IFA011 V0.7.3"
        required: false
      l3_address_data:
        # Shall be present when the addressType is ip_address.
        description: >-
          Provides the information on the IP addresses to be assigned to the connection point
          instantiated from the parent Connection Point Descriptor.
        type: tosca.datatypes.nfv.L3AddressData
        required: false

  tosca.datatypes.nfv.VirtualNetworkInterfaceRequirements:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0-csd04]'
      citation_location: 5.3.4
    description: >-
      The VirtualNetworkInterfaceRequirements type is a complex TOSCA data type used to describe
      VirtualNetworkInterfaceRequirements information element as defined in [ETSI GS NFV-IFA 011],
      it provides the information to specify requirements on a virtual network interface realising
      the CPs instantiated from this CPD.
    derived_from: tosca.datatypes.Root
    properties:
      name:
        description: >-
          Provides a human readable name for the requirement.
        type: string
        required: false
      description:
        description: >-
          Provides a human readable description for the requirement.
        type: string
        required: false
      support_mandatory:
        description: >-
          Indicates whether fulfilling the constraint is mandatory (TRUE) for successful operation
          or desirable (FALSE).
        type: boolean
        required: false
      requirement:
        description: >-
          Specifies a requirement such as the support of SR-IOV, a particular data plane
          acceleration library, an API to be exposed by a NIC, etc.
        type: string # ERRATUM: the spec says "not specified", but TOSCA requires a type
        required: true

  tosca.datatypes.nfv.ConnectivityType:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0-csd04]'
      citation_location: 5.3.5
    description: >-
      The TOSCA ConnectivityType type is a complex TOSCA data type used to describe ConnectivityType
      information element as defined in [ETSI GS NFV-IFA 011].
    derived_from: tosca.datatypes.Root
    properties:
      layer_protocol:
        description: >-
          Identifies the protocol this VL gives access to (ethernet, mpls, odu2, ipv4, ipv6,
          pseudo_wire).
        type: string
        required: true
        constraints:
        - valid_values: [ ethernet, mpls, odu2, ipv4, ipv6, pseudo_wire ]
      flow_pattern:
        description: >-
          Identifies the flow pattern of the connectivity (Line, Tree, Mesh).
        type: string
        required: false

  tosca.datatypes.nfv.RequestedAdditionalCapability:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0-csd04]'
      citation_location: 5.3.6
    description: >-
      RequestAdditionalCapability describes additional capability for a particular VDU.
    derived_from: tosca.datatypes.Root
    properties:
      request_additional_capability_name:
        description: >-
          Identifies a requested additional capability for the VDU.
        type: string
        required: true
      support_mandatory:
        description: >-
          Indicates whether the requested additional capability is mandatory for successful
          operation.
        type: string
        required: true
      min_requested_additional_capability_version:
        description: >-
          Identifies the minimum version of the requested additional capability.
        type: string
        required: false
      preferred_requested_additional_capability_version:
        description: >-
          Identifies the preferred version of the requested additional capability.
        type: string
        required: false
      target_performance_parameters:
        description: >-
          Identifies specific attributes, dependent on the requested additional capability type.
        type: map
        entry_schema:
          type: string
        required: true

  tosca.datatypes.nfv.VirtualMemory:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0-csd04]'
      citation_location: 5.3.7
    description: >-
      VirtualMemory describes virtual memory for a particular VDU.
    derived_from: tosca.datatypes.Root
    properties:
      virtual_mem_size:
        description: Amount of virtual memory.
        type: scalar-unit.size
        required: true
      virtual_mem_oversubscription_policy:
        description: >-
          The memory core oversubscription policy in terms of virtual memory to physical memory on
          the platform. The cardinality can be 0 during the allocation request, if no particular
          value is requested.
        type: string
        required: false
      numa_enabled:
        description: >-
          It specifies the memory allocation to be cognisant of the relevant process/core
          allocation. The cardinality can be 0 during the allocation request, if no particular value
          is requested.
        type: boolean
        required: false

  tosca.datatypes.nfv.VirtualCpu:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0-csd04]'
      citation_location: 5.3.8
    description: >-
      VirtualMemory describes virtual memory for a particular VDU.
    derived_from: tosca.datatypes.Root
    properties:
      cpu_architecture:
        description: >-
          CPU architecture type. Examples are x86, ARM.
        type: string
        required: false
      num_virtual_cpu:
        description: >-
          Number of virtual CPUs.
        type: integer
        required: true
      virtual_cpu_clock:
        description: >-
          Minimum virtual CPU clock rate.
        type: scalar-unit.frequency
        required: false
      virtual_cpu_oversubscription_policy:
        description: >-
          CPU core oversubscription policy.
        type: string
        required: false
      virtual_cpu_pinning:
        description: >-
          The virtual CPU pinning configuration for the virtualized compute resource.
        type: tosca.datatypes.nfv.VirtualCpuPinning
        required: false

  tosca.datatypes.nfv.VirtualCpuPinning:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0-csd04]'
      citation_location: 5.3.9
    description: >-
      VirtualCpuPinning describes CPU pinning configuration for a particular CPU.
    derived_from: tosca.datatypes.Root
    properties:
      cpu_pinning_policy:
        description: >-
          Indicates the policy for CPU pinning.
        type: string
        constraints:
        - valid_values: [ static, dynamic ]
        required: false
      cpu_pinning_map:
        description: >-
          If cpuPinningPolicy is defined as "static", the cpuPinningMap provides the map of pinning
          virtual CPU cores to physical CPU cores/threads.
        type: map
        entry_schema:
          type: string
        required: false

  tosca.datatypes.nfv.VnfcConfigurableProperties:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0-csd04]'
      citation_location: 5.3.10
    # ERRATUM: description is mangled in spec
    description: >-
      VnfcConfigurableProperties describes additional configurable properties of a VNFC.
    derived_from: tosca.datatypes.Root
    properties:
      additional_vnfc_configurable_properties:
        description: >-
          Describes additional configuration for VNFC.
        type: map
        entry_schema:
          type: string
        required: false
`
}
