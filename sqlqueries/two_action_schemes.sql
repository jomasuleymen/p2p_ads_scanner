WITH common_table as (
    SELECT advertisements.*,
        apm.payment_method_id
    FROM "advertisements"
        INNER JOIN ads_payment_methods AS apm ON apm.advertisement_id = advertisements.id
        INNER JOIN advertisers ON advertisers.id = advertisements.advertiser_id
    WHERE asset_id = 1
        AND fiat_id = 1
),
buy_data as (
    SELECT *
    FROM "common_table"
    WHERE trade_side = 'SELL'
),
sell_data as (
    SELECT *
    FROM (
            SELECT *,
                ROW_NUMBER() OVER (
                    PARTITION BY exchange_id,
                    payment_method_id
                    ORDER BY price ASC
                ) as ranked_order
            FROM "common_table"
            WHERE trade_side = 'BUY'
        ) as mt
    WHERE mt.ranked_order = 1
        and price BETWEEN (
            SELECT Min(price)
            FROM "buy_data"
        ) * 1.001
        AND (
            SELECT Max(price)
            FROM "buy_data"
        ) * 1.069200
),
buy_data_f as (
    SELECT *
    FROM "buy_data"
    WHERE price BETWEEN (
            SELECT Min(price)
            FROM "sell_data"
        ) / 1.069200
        AND (
            SELECT Max(price)
            FROM "sell_data"
        ) / 1.001
    ORDER BY price ASC
    LIMIT 15
), sell_data_f as (
    SELECT *
    FROM "sell_data"
    WHERE price BETWEEN (
            SELECT Min(price)
            FROM "buy_data_f"
        ) * 1.001
        AND (
            SELECT Max(price)
            FROM "buy_data_f"
        ) * 1.069200
)
SELECT buy.id as buy_id,
    Round((sell.price / buy.price - 1) * 100, 3) AS spread_percent,
    sell.id as sell_id
FROM buy_data_f buy
    INNER JOIN sell_data_f sell ON sell.price / buy.price BETWEEN 1.001 AND 1.069200
ORDER BY spread_percent desc
LIMIT 15