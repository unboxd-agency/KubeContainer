# The Measurement Standard

A kube is a box with meaning and a measurement — a bounded thing whose
value is *defined* (what it is) and *measured* (how much). This document
is the constitutional rule for how meaning and measurement are
expressed, so that value on the fabric is never a private invention
dressed as a number.

## The standards rule

Meaning and measurement are presented in terms of real-world,
international, globally accepted standards — always, by default, and
named. A kube's quantities ride recognized units and schemas, never
bespoke ones:

- *Money* — ISO 4217 currency codes; amounts as decimal minor units,
  never floats.
- *Time* — ISO 8601 / RFC 3339 timestamps; UTC at the record, with
  zone where it carries meaning (the temporal dimension).
- *Place* — ISO 3166 regions; Open Location Code for points (the
  geospatial dimension).
- *Identity* — W3C DID / Verifiable Credentials (did:web).
- *Goods & trade* — GS1 (GTIN/GLN), HS commodity codes, UN/CEFACT
  where applicable.
- *Units* — SI / ISO 80000 for physical quantity.
- *Language & domain* — published schemas (ISO, W3C, the relevant SDO)
  for each bounded context; meaning cites its ontology.

The rule generalizes: where a globally accepted standard exists for a
quantity, the kube uses it and names it — measurement without a named
standard is an opinion with a number on it.

## The gap clause (where no standard exists)

In the absence of any globally accepted standard for a given meaning or
measurement, the definition and delivery of it is left to the contract
holder. The contract holder then:

1. *Defines it explicitly* — the unit, the method, the boundaries —
   published in the registry entry, not assumed.
2. *Delivers against that definition* — the world-test measures what
   was defined; the verdict is against the holder's own published
   meter.
3. *Owns the consequence* — a holder-defined measure is the holder's
   to stand behind, and is marked as holder-defined, never passed off
   as a global standard it is not.

Holder-defined is honest; holder-defined-disguised-as-standard is the
cardinal measurement sin (a claim wearing a verdict's clothes).

## The supply-chain integrity rule

A kube's bill of materials (BOM) and supply-chain record (SCM) must
contain no prior transaction before the end-user touch point. The
measurement and its provenance begin clean at delivery to the
principal: what is measured and sold is *this* kube's value, established
at *this* touch, not inherited, laundered, or back-dated from upstream
dealings. Provenance runs forward from a clean origin; it is not a
ledger of everyone who handled the goods before they became the
customer's. (This protects the customer — they buy a clean, attributable
unit — and the standard — no upstream noise contaminates the measure.)

## The two paths to a standard you can keep

When no global standard exists and a holder-defined measure is not
enough — when the value must be *interoperable*, comparable across
parties — there are exactly two honest paths, and no third:

1. *Own the entire supply chain end to end* — control every step so
   the standard is yours, your way, internally consistent because one
   party defines and delivers all of it. (Vertical integration as the
   price of a private standard.)
2. *Reach out / form an international body* — convene the parties and
   align on a common, interoperable way of defining and measuring the
   quantity, so the standard is shared and no single party owns it.
   (The SDO path — slower, but the only route to a measure others will
   accept as theirs too.)

The dishonest third path — declaring a private measure "the standard"
without owning the chain or convening the body — is forbidden: it is the
openwashing of measurement, a claim of universality with neither the
control nor the consent that would make it true.

## Worked example: the chocolate box

A box of chocolates is the kube in miniature, and its value composes
transparently from measured parts — exactly how a kube's price is built:

```
chocolate box value
  = Σ (value of each chocolate)     # the contents, the work itself
  + service charges                 # the boxing, the keeping, the handling
  + taxes                           # the jurisdiction's claim, by its rate
```

Each line rides a standard: the chocolates' value in ISO-4217 currency,
minor units; the service charge an itemized, named line (not a hidden
markup — priced by the piece); the tax computed at the rate of the ISO
3166 jurisdiction of the touch point, by that jurisdiction's published
code. The bill is itemized to the same grain as the evidence — the buyer
sees what the contents cost, what the service cost, what the state took —
and nothing in the total is a number without a named basis. Generalize
the three lines and you have every kube's price: *contents + service +
obligations*, each measured to standard, each on the manifest, summing
to a total anyone can re-derive. (And per the supply-chain rule: the
chocolates' value begins at the customer's touch point — the wholesale
dealings that put them in the box are not the customer's BOM.)

## The transaction formula (canonical)

At every transaction the value is computed fresh, never accumulated:

```
transaction value
  = quantity × unit value           # the contents: qty x value, to standard
  + service charge (current platform)  # levied by the platform operating THIS transaction
  + taxes                           # at the current jurisdiction's published rate
```

The load-bearing word is *current*. The service charge is the one
levied by the platform operating *this* transaction — not a sum of every
platform that ever touched the goods upstream — and the tax is the
current touch point's jurisdiction at its current rate. This is the
clean-BOM rule expressed as arithmetic: each transaction begins fresh at
`quantity × unit value` and adds only what *this* touch legitimately
adds (the current operator's service, the current jurisdiction's tax).
No prior transaction's charges or taxes ride forward; value is not a
sediment of every hand it passed through but a clean computation at each
touch. The chocolate box is the special case (quantity = the chocolates,
service = the boxing, taxes = the jurisdiction); the formula is the
general law, and it holds at every transfer on the fabric — buy, resell,
sublease, settle — each a fresh, itemized, standard-denominated total
anyone can re-derive from its three named lines.

And the entire calculation is available at every point of the supply
chain — transparency is not the buyer's privilege alone but the chain's
property. At any transfer, any participant can see the full breakdown of
the value at their point: the quantity and unit value, the service
charge the current platform levied and why, the tax and its
jurisdiction — and the running composition up to (but never before) the
clean origin of their own touch. No black-box totals, no "trust the
price"; the arithmetic that produced the number is inspectable wherever
the number travels, the same way provenance travels with the kube. A
price whose computation cannot be seen at the point it is charged is a
claim, not a measure — and the fabric carries measures.

And the visible arithmetic carries a visible incentive: more stops, more
cost. Every transfer that adds a platform adds that platform's service
charge and its jurisdiction's tax — so a chain with more intermediaries
costs more, transparently, line by line, and the buyer sees exactly how
many hands took a cut and what each took. This is not a defect to hide
but a signal to expose: where the old supply chain buried its markups so
no one could count the middlemen, the fabric itemizes every stop, which
makes disintermediation a measurable choice rather than a marketing
claim. Fewer stops is cheaper and the bill proves it; a direct touch
beats a brokered one and the arithmetic shows by how much. The platform
that adds a stop must justify its charge in the open, at the point it
levies it — and a stop that adds cost without adding value is now
visible to the one party with reason to remove it.

So the fabric does not dictate the chain; it lets the people find it.
With every stop's cost itemized and every route's total re-derivable,
participants discover the optimum supply chain themselves — the ideal
number of players and the points of stops between source and
destination — by reading the arithmetic, not by trusting a broker.
The platform's job is to make the costs legible; the market's job is to
choose the path; and because the costs are true and visible, the path
the people converge on is the genuinely optimal one, not the one a
hidden intermediary preferred. This is finding the critical path, and
finding it is a real-world optimization problem — minimize total cost
and total latency across the graph of possible routes, subject to the
real constraints (which players can actually do the work, which
stops are required by law or geography, which edges the contracts
permit). The fabric supplies what the optimization needs: a graph that
resolves (every player and route a node and edge), true edge weights
(the itemized costs and times, no hidden tolls), and verdicts (each
route's claims checkable). Optimization over a graph of true weights is
a solved class of problem the moment the weights stop lying — and making
the weights stop lying is exactly what the measurement standard does.

## Why this is constitutional

Measurement is where value becomes checkable, and value is what the
fabric sells. A measure on a recognized standard is a verdict anyone can
re-run; a measure on a private unit is a claim only the seller can
confirm — which collapses provenance, assurance, and the registry's
whole worth. The kube carries meaning and measurement; this standard
keeps both honest, so that "priced by the piece" means priced in units
the world already agreed on, or units the holder openly owns, and never
units pretending to be either.
