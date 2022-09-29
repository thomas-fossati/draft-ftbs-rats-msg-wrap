---
v: 3

title: "RATS Conceptual Messages Wrapper"
abbrev: "RATS CMW"
docname: draft-ftbs-rats-msg-wrap-latest
category: std

ipr: trust200902
area: "Security"
workgroup: "Remote ATtestation ProcedureS"
keyword: Internet-Draft

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

informative:


--- abstract

TODO Abstract


--- middle

# Introduction

TODO Introduction


# Conventions and Definitions

{::boilerplate bcp14-tagged}


# CDDL

~~~ cddl
{::include cddl/cmw.cddl}
~~~

## Examples

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

TODO IANA



--- back

# Acknowledgments
{:numbered="false"}

TODO acknowledge.
