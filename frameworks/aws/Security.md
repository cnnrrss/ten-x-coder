# Security


### Redshift

- Users _inbound_ access to an Amazon Redshift cluster can be managed through **security groups**

- **Access Management** to Redshift cluster can be managed through **IAM**


### Encryption in-transit

- Enable ssl
    - Encryption of connection between client and cluster (1st layer)
- SSL CA Certificate
    - SSL certificate pem key to connect to the cluster

### Encryption at-rest

- Two options: 
    - Key Management Service (KMS)
    - Hardware Security Module (CloudHSM and On-Prem HSM)
- When enabled, encryption of: data blocks, system metadata, and snapshots
- Encryption on a Redshift cluster is _immutable_
    - You'll have to unload data and reload in a new cluster

#### KMS

Four-tier hierarchy of encryption keys:
    1) Master key
    2) Cluster encryption key
    3) Database encryption key
    4) Data Encryption keys

#### HSM
Physical devices that safeguard and manage digital keys for strong authentication and provide cryptoprocessing.

Contractual, regulatory requirements may determine if a HSM should be used.