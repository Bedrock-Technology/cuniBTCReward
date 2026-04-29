create schema cunibtcreward collate utf8mb4_unicode_ci;
use cunibtcreward;

create table cursors
(
    id           bigint unsigned auto_increment
        primary key,
    created_at   datetime(3)                 null,
    updated_at   datetime(3)                 null,
    deleted_at   datetime(3)                 null,
    chain_id     bigint unsigned default '0' not null,
    chain_name   varchar(255)                null,
    block_number bigint unsigned default '0' null,
    constraint t_chainid_name
        unique (chain_id, chain_name)
);

create index idx_cursors_deleted_at
    on cursors (deleted_at);

create table evm_transactions
(
    id              bigint unsigned auto_increment
        primary key,
    created_at      datetime(3)                 null,
    updated_at      datetime(3)                 null,
    deleted_at      datetime(3)                 null,
    address         varchar(255)                null,
    chain_id        bigint unsigned default '0' not null,
    hash            varchar(255)                null,
    block_number    bigint unsigned default '0' null,
    block_timestamp bigint unsigned default '0' null,
    amount          decimal(38)     default 0   null,
    log_method      varchar(255)                null,
    memo            varchar(255)                null,
    constraint t_address_chainid_hash
        unique (address, chain_id, hash)
);

create index idx_evm_transactions_deleted_at
    on evm_transactions (deleted_at);

create index t_blocknumber
    on evm_transactions (block_number);

create index t_blocktimestamp
    on evm_transactions (block_timestamp);

