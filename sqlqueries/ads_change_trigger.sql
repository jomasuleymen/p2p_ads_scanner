create table IF NOT exists public.ads_change(
	advertiser_id int8 NOT NULL,
	exchange_id int8 not null,
	fiat_id int8 not NULL,
	asset_id int8 not NULL,
	trade_side bpchar(4) not NULL,
	price float4 not NULL,
	changed_asset_num float4 not null,
	payment_methods text [],
	updated int8 NULL,
	CONSTRAINT fk_ads_change_advertiser FOREIGN KEY (advertiser_id) REFERENCES public.advertisers(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX if not exists idx_ads_change_advertiser_id ON public.ads_change USING btree (advertiser_id);

CREATE or replace FUNCTION ads_change_func() 
   RETURNS TRIGGER 
   LANGUAGE PLPGSQL
AS $$
declare 
	changed_asset_num ads_change.changed_asset_num%type := 0;
BEGIN
   
	changed_asset_num := ABS(new.available_asset_num - old.available_asset_num);
	
	insert into ads_change values
	(new.advertiser_id, new.exchange_id, new.fiat_id, new.asset_id,
		new.trade_side, new.price, changed_asset_num,
		new.payment_methods_list, new.updated);
	
    RETURN NEW;
END;
$$;

create or replace trigger ads_update_trigger
	AFTER UPDATE ON advertisements
    FOR EACH ROW
    WHEN (OLD.available_asset_num <> NEW.available_asset_num)
    execute function ads_change_func();