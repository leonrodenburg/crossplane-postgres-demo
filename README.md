# E2E Live - Crossplane demo

Repository showing how you can directly provision cloud resources with Crossplane. It shows three ways of deploying:

1. Directly creating managed resources (`rds-instance.yaml`)
2. Deploying managed resources through Helm and connecting to an app (`./chart`)
3. Using composition to abstract away cloud-specific implementation (`./composition`)
