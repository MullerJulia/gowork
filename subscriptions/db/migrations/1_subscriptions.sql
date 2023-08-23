-- migrate:up
CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY, 
    user_id VARCHAR(256),  
    subscription_id VARCHAR(256), 
    status VARCHAR(256), 
    charge_amount INT
);

-- migrate:down
drop table subscriptions;