# The Apache Attic — a Mining Review

Status: assessment, 2026-06-12. The founder's direction: the platform
backend is Java because the Apache ecosystem is the toolbox — and the
Attic (https://attic.apache.org/) is the part of that toolbox whose
keepers have left. Retired Apache projects are explicitly forkable;
the Attic rule (deploy/AGENT-STACK.md) governs every adoption: enter
only pinned at the last release, with a keeper named, and nothing
counts as adopted until it passes this house's gauntlet — adoption is
tested, not declared.

## The shelves worth mining

| Project (retired) | What it is | The seat it could fill | Verdict |
|---|---|---|---|
| Portals: Jetspeed + Pluto (2025) | Java enterprise portal + the reference portlet container | the panel/portal layer — Liferay's open cousin, freshly retired, code still current-era Java | mine first: newest retirement, largest overlap with the platform backend ruling |
| Marmotta (2018) | linked-data platform: JSON-LD, SPARQL, LDP server in Java | the record graph's server — `eval/graph.jsonld` already speaks its language | mine second: direct fit with RecordGraph/SchemaKeeper; age means a dependency audit first |
| ODE (2023) | BPEL workflow/orchestration engine | the flow engine seat (the Temporal/Orkes office) | study only: BPEL is a dated dialect; mine the patterns, not the engine |
| Wave (2018) | real-time collaborative editing (Google Wave) | the collaboration space's ancestor | study only: concepts for Agent-Space's room, code too old |
| Lenya / Cocoon | XML CMS / web framework | content backends | pass: superseded by living headless backends on the contract |
| Oltu, Shindig, Stanbol | OAuth lib / gadgets / semantic content | identity, widgets, semantics | pass: living standards replaced them |

## AutonomyX

The founder's naming, 2026-06-13: AutonomyX is Apache — the platform
brand under which Attic adoptions are revived and kept, its door
openautonomyx.com (the arithmetic platform at
platform.openautonomyx.com, the backend per the backend contract).
Every AutonomyX component is an Attic adoption under the law below,
or a living Apache dependency, or this house's own tool.

## The full inventory (the Attic, 2026-06)

abdera, ace, any23, apex, archiva, aurora, avalon, axis-sandesha-c,
axis-savan-c, axis-savan-java, axkit, bahir, beehive, bloodhound,
buildr, chemistry, chukwa, clerezza, click, climate, cocoon,
continuum, crimson, crunch, deltacloud, devicemap, directmemory,
drat, eagle, esme, etch, excalibur, falcon, forrest, giraph, gora,
griffin, hama, harmony, hawq, hivemind, ibatis, jakarta-cactus,
jakarta-ecs, jakarta-oro, jakarta-regexp, jakarta-slide,
jakarta-taglibs, jakarta, jclouds, joshua, juddi, kibble, labs,
lens, lenya, lucy, marmotta, mesos, metamodel, metron, mnemonic,
mrunit, muse, mxnet, ode, ojb, olingo, oltu, onami, oodt, oozie,
pivot, polygene, portals, predictionio, quetzalcoatl, rave, reef,
river, sentry, servicemix, shale, shindig, sqoop, stanbol, stdcxx,
stratos, streams, submarine, tajo, tiles, trafficcontrol, trafodion,
tuscany, twill, usergrid, vxquery, whirr, wink

## Additional shelves flagged on full review

| Project | What it is | The seat |
|---|---|---|
| jclouds | the Java multi-cloud toolkit | the founder's multi/hybrid-cloud dictation, pre-built |
| mesos | the cluster manager | study: the pre-Kubernetes lesson, not a substrate |
| stratos | a Java PaaS | study: platform patterns for AutonomyX |
| usergrid | a Java backend-as-a-service | backend-contract candidate (heavy: Cassandra behind it) |
| juddi | UDDI service registry | the agent-registry ancestor; patterns for AgentRegistry |
| rave + shindig | web widget/portal mashups | the panel's widget layer, with portals/jetspeed |
| oozie, falcon, ode | workflow engines | the flow office's Java ancestors; patterns only |

## The law applied

- Pinned: an Attic adoption names the exact final release and its
  checksum before a line is touched.
- Kept: the keeper is named in the tool registry the day the fork is
  made — an empty seat is taken, never squatted.
- Tested: the founder's word, "we need to test" — an adopted codebase
  enters the gauntlet (compile, vet, lint where applicable, its own
  test suite resurrected) and earns its registry row with a green
  verdict, not a citation.
- Licensed: Attic code is Apache-2.0; it may be *used* as a
  dependency, but this house's license of record is unchanged and the
  divergence is named per adoption.
