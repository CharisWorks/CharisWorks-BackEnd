# Database Design

## users
- id
    - PK, Firebase
- stripe_account_id
    - UK, Stripe
    - null ok
- history_user_id
    - UK, FK, int
- created_at
    - timestamp

## history_users
- id
    - PK, SK, int
- user_id
    - FK, Firebase
- real_name
- display_name
- description
    - text
    - null ok
- created_at
    - timestamp



## shippings
- id
    - PK, FK, Firebase
- zip_code
- address_1
- address_2
- address_3
    - null ok
- phone_number



## items
- id
    - PK, SK, int
- manufacturer_user_id
    - FK, Firebase
- history_item_id
    - UK, FK, int

## history_items
- id
    - PK, SK, int
- item_id
    - FK, int
- name
- price
    - int
- status
- stock
    - int
- size
    - int
    - null ok
- description
    - text
    - null ok
- tags
    - null ok



## transactions
- id
    - PK, SK, int
- manufacturer_user_id
    - FK, Firebase
- purchaser_user_id
    - FK, Firebase
- item_id
    - FK, int
- quantity
    - int
- tracking_id
- created_at
    - timestamp
- zip_code
- address
- phone_number
- history_manufacturer_user_id
    - FK, int, (履歴保持用)
- history_purchaser_user_id
    - FK, int, (履歴保持用)
- history_item_id
    - FK, int, (履歴保持用)



## carts
- id
    - PK, SK, int
- purchaser_user_id
    - FK, Firebase
- item_id
    - FK, int
- quantity
    - int
