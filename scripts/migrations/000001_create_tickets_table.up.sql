create table tickets
(
    id          uuid                                      not null,
    title       text                                      not null comment 'title',
    description text                                      null,
    status      tinyint      default 0                    not null comment 'status of task 0:UNSPECIFIED 1:TODO 2:INPROGRESS 3:DONE',
    created_at  timestamp(6) default current_timestamp(6) not null comment 'created time',
    updated_at  timestamp(6) default current_timestamp(6) not null on update current_timestamp(6) comment 'updated time',
    primary key (id)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_unicode_ci comment 'ticket list';

create index idx_tickets_status on tickets (status);
create index idx_tickets_created_at on tickets (created_at);
create index idx_tickets_updated_at on tickets (updated_at);
