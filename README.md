# coop

Estellencs.coop resource management

## Configuration

Coop if configured with `coop.yaml` file.
It should be located at `$COOP_HOME` or `$HOME/coop`.

Config file example:
```yaml
database:
  name: /Users/username/coop/database/coop.db
pdf:
  deliveryNote: "DeliveryNote-%s.pdf"
```

## Generate delivery note PDF and update status

- Generate PDFs for all `draft` delivery notes: `coop pdf dn` and update status to `published`.
- Generate PDF for the delivery note 666: `coop pdf dn --id=666` and update status to `published`.