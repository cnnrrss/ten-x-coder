# TCO and Cloud Economics

## Mod 1: What is TCO?
### Key Concepts
**TCO** = initial purchase price of an asset + ops costs
    - Look beyond purchase cost
    - Consider long-term cost
    - Lower TCO will provide higher value over time

**Timeboxing** = set period of time that factors into a TCO calculation and analysis

**Stakeholders** = TCO of products and services is important to both business and IT stakeholders

**TCO Discussions** = Articulating various aspects of TCO can enable effective business conversations

TCO Example: How much does a car really cost?
- Initial Purchase Price (MSRP)
- Registration & Insurance cost (ongoing)
- Fuel (season, model, engine type)
- Maintenance (regardless of new or used)

TCO = Not just purchase price but expenses buyer may incur throughout a product's lifespan. (Used car might breakdown and have huge maintenance cost, new car comes with 3 year warranty)

#### TCO and AWS
**Challenges** = TCO-related IT challenges are a major business issue enterprises face

**Foundation** = Building IT costs baseline for all aspects of data center sets proper foundation for demonstrating solution benefit.

**AWS TCO** = TCO for AWS models total costs of ownership for acquiring & running IT infra env.

**Compare the costs** AWS vs on-prem vs colocation
- Create a detailed cost comparison
- Model a specific workload and compare it
- Communicate value to customers

**Quiz** When looking at a product / solution TCO, you must take into account: (Choose all that apply)
- Set period of time of use
- ~~Economies of scale~~
- Purchase price
- Operating Costs

#### How is TCO Lower with AWS
- Replace up-front cap exp with low var cost
- Utilizes economies of scale to continually lower costs
- Pricing model choice to support var and stable workloads
- Delivers greater savings as they grow

### TCO Components
1) Servers, Network, Hardware
    - Compute Instances(EC2) vs Customer Servers
2) Operating System & Virtualization Software
    - Windows & Linux os and virtualization software costs are included in AWS
3) Data Center Colocation & Floor Space
    - Floor space to purchase or lease real estate can be major expense
4) Power & Cooling
    - Industry averages on cost of utilities for consistent temp across data center are included as cost assumptions in AWS TCO analysis
5) Data Center Personnel
    - Modern data centers employ automation, but still need administrators
6) Storage Redundancy
    - Durability (data saved regardless e.g. natural disaster)
    - Availability (info is immediately accessible)
    - S3 offers eleven nines of durability
7) Resource Management & Software Automation
    - CloudFormation = Create and manage a collection of related AWS resources
        - No need to figure out the order for provisioning AWS services
        - Higher productivity, lower TCO
8) Software-Defined Networking
    - Software-defined networking via Amazon - VPC
    - One person can manage complex env

#### Conceptual representation of TCO
On-Premises = Customer buys all TCO components
Colocation = Buy most of TCO components, colocation bundles power/floor space together
AWS = provides all TCO components & physical security

#### TCO Data collection Template
Tool for collaborating with customers to simplify TCO data collection

**Quiz** What are 3 cost components of AWS TCO analysis?
- ~~Cloud, colocation, on-perm~~
- Data center hardware, software, power costs
- ~~Network resiliency, server uptime, & floor space~~
- ~~Availability, durability, and scalability~~

### AWS Pricing Model

#### AWS Pricing Model
On-Premises = Opex+++, Capex+++
Colocation = Opex++, Capex++
AWS = Opex+

#### AWS Pricing Philosophy
- More customers on AWS, cheaper the price (economies of scale)
- Cost savings passed to customers
- AWS reduces pricing all the time

#### Pricing Characteristics

3 greatest influences on cost:

- **Data transfer**: pricing can flux, data in = free, data out is min charge (varies on market), data out = tiered pricing

- **Storage**: Pricing for storage is dependent on customer's needs and usage (may recieve discounts)

- **Compute**: EC2 (On-Demand, Reserved, Spot, Dedicated), Lambda
    - On-Demand: compute capacity by the hour
    - Reserved Instance: significant discounts (75% to on-demand), All Upfront, Partial Upfront, No Upfront
    - Spot Instances: bid for unused Amazon EC2 capacity (Spot price is set by the market), supply and demand, place a detailed request.
    - Dedicated: Single instance not dedicated for other tasks. Hourly per-instance usage fee and dedicated per-region fee

**Quiz** What instance type is good if customer wants to bid on current market price of compute power?
    - ~~On-Demand~~
    - Spot
    - ~~Reserved~~
    - ~~Upfront~~

## Mod 2: Target Customers

## Mod 3: Discussing AWS Economics

## Mod 4: Data Center Economics

## Mod 5: Delivering the TCO Message

## Mod 6: TCO Modeling

## Mod 7: Presenting the results of TCO

## Mod 8: Addressing Common Issues & Objections

## Accreditation