WITH common_table as (
    SELECT advertisements.*,
        apm.*
    FROM "advertisements"
        INNER JOIN ads_payment_methods AS apm ON apm.advertisement_id = advertisements.id
    WHERE fiat_id = 1
),
base_quote_table as (
    select ads.*,
        bt.quote_id as exchange_to,
        bt.bid_price as spot_price,
        ads.price / bt.bid_price AS new_price,
        from common_table ads
        join book_tickers bt on bt.exchange_id = ads.exchange_id
        and (ads.asset_id = bt.base_id)
),
quote_base_table as (
    select ads.*,
        bt.base as exchange_to,
        bt.ask_price as spot_price,
        ads.price * bt.ask_price AS new_price,
        from common_table ads
        join book_tickers bt on bt.exchange_id = ads.exchange_id
        and (ads.asset_id = bt.quote_id)
),
all_ads as (
    select *
    from quote_base_table
    union all
    select *
    from base_quote_table
),
sell_side AS (
    select *
    FROM common_table
    WHERE trade_side = 'SELL'
),
buy_side AS (
    SELECT *
    FROM all_ads
    WHERE trade_side = 'BUY'
)
select buy_side.id as buy_id,
    buy_side.asset_id as buy_asset,
    buy_side.exchange_id as buy_exchange,
    buy_side.price as buy_price,
    buy_side.spot_price as spot_price,
    (sell_side.price / buy_side.new_price - 1) * 100 AS spread_percent,
    sell_side.id as sell_id,
    sell_side.asset_id as sell_asset,
    sell_side.exchange_id as sell_exchange,
    sell_side.price as sell_price,
from buy_side
    join sell_side on buy_side.exchange_to = sell_side.asset_id
where sell_side.price / buy_side.new_price BETWEEN 1.0 + 0 / 100.0 AND 1.0 + 150 / 100.0
order by spread_percent desc
limit 15