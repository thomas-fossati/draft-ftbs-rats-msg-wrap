---
v: 3

title: "RATS Conceptual Messages Wrapper"
abbrev: "RATS CMW"
docname: draft-ftbs-rats-msg-wrap-latest
category: std
consensus: true
submissionType: IETF

ipr: trust200902
area: "Security"
workgroup: "Remote ATtestation ProcedureS"
keyword: [ evidence, attestation result, endorsement, reference value ]

stand_alone: yes
smart_quotes: no
pi: [toc, sortrefs, symrefs]

author:
 - name: Hannes Tschofenig
   organization: arm
   email: hannes.tschofenig@arm.com
 - name: Henk Birkolz
   organization: Fraunhofer SIT
   email: henk.birkholz@sit.fraunhofer.de
 - name: Ned Smith
   organization: Intel
   email: ned.smith@intel.com
 - name: Thomas Fossati
   organization: arm
   email: thomas.fossati@arm.com

normative:
  RFC4648: base64
  RFC6838: media-types
  RFC7252: coap
  RFC8259: json
  RFC8610: cddl
  RFC9165: cddlplus
  RFC9277:
  STD94:
    -: cbor
    =: RFC8949

informative:
  I-D.ietf-rats-architecture: rats-arch

--- abstract

This document defines two encapsulation formats for RATS conceptual
messages.

--- middle

# Introduction

This document defines two encapsulation formats for RATS conceptual
messages ({{Section 8 of -rats-arch}}).

TODO use cases (TLS, cert extensions, long-term storage, use in API
payloads, other)

# Conventions and Definitions

{::boilerplate bcp14-tagged}

In this document, CDDL {{-cddl}} {{-cddlplus}} is used to describe the
data formats.

The reader is assumed to be familiar with the vocabulary and concepts
defined in {{-rats-arch}}.

# Conceptual Message Wrapper Encodings

Two different RATS Conceptual Message Wrapper (CMW) data types are
defined: one using a CBOR or JSON array ({{type-n-val}}), another using
CBOR tags ({{cbor-tag}}).

## Type and Value Array {#type-n-val}

The CMW described in {{fig-cddl}} is a CDDL array of two elements:

* type: a string encoding a media type {{-media-types}} or an
  unsigned integer corresponding to a CoAP Content-Format {{-coap}}

* value: the RATS conceptual message serialisation according to the
  specified type.

This kind of CMW can be encoded as CBOR {{-cbor}} or JSON {{-json}}.

When using JSON, the value field is encoded as Base64 using the URL and
filename safe alphabet (Section 5 of {{-base64}}) without padding.

When using CBOR, the value field is serialized as a CBOR bytes string.

~~~ cddl
{::include cddl/cmw.cddl}
~~~
{: #fig-cddl artwork-align="left"
   title="CDDL definition"}

## CBOR Tags derived from CoAP Content Formats {#cbor-tag}

If a CoAP Content Format exists for the RATS conceptual message, the
TN() transform defined in {{Appendix B of RFC9277}} can be used to
derive a CBOR tag in range [1668546817, 1668612095].

# Examples

The (equivalent) examples below assume the media type
`application/vnd.example.rats-conceptual-msg` has been registered
alongside a corresponding CoAP content format `30001`.  The CBOR tag
`1668576818` is derived applying the TN transform as described in
{{cbor-tag}}.

~~~ cbor-diag
{::include cddl/example-cbor-1.diag}
~~~
{: #fig-example-cbor artwork-align="left"
   title="CBOR encoding"}

~~~ cbor-diag
{::include cddl/example-json-1.diag}
~~~
{: #fig-example-json artwork-align="left"
   title="JSON encoding"}

~~~ cbor-diag
1668576818(h'abcdabcd')
~~~
{: #fig-example-cbor-tag artwork-align="left"
   title="CBOR tag"}

# Security Considerations

TODO Security


# IANA Considerations

When registering a new media type for evidence, in addition to its
syntactical description, the author SHOULD provide a public and stable
description of the signing and verification procedures associated with
the data format.



--- back

# Acknowledgments
{:numbered="false"}

TODO acknowledge.

- vim: tw=72 ts=4 sw=4 et