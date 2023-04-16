create table if not exists users
(
    id         bigint       not null,
    username   varchar(255) not null,
    password   varchar(255) not null,
    token      text         null,
    created_at timestamp(6) not null default current_timestamp(6),
    updated_at timestamp(6) not null default current_timestamp(6) on update current_timestamp(6),
    constraint pk_users primary key (id)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment = 'users table';