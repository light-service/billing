# billing

## Resource

product(id,name)
resource(id,key,name,record_unit,display_unit,display_unit_conversion,display_unit_price,product_id)

## Consumption

resource_pack_template(name,resource_key,resource_id,amount,fixed)
resource_pack
resource_pack_detail
consumption(bill_key, accrual, resource_key, resource_id, by_resource_pack)

## Order

discount
coupon_template
coupon
coupon_detail
order()

## Accounting

account(account_key,balance_type,can_negative)
journal(id,name,debit_account,credit_account)
journal_resource(journal_id,resource_key)
journal_entry()