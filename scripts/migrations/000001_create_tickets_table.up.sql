create table if not exists tickets
(
    id         bigint unsigned                        not null,
    title      text                                   not null comment 'title',
    status     tinyint      default 0                 NOT NULL COMMENT 'status of task 0:UNSPECIFIED 1:TODO 2:INPROGRESS 3:DONE',
    created_at timestamp(6) default CURRENT_TIMESTAMP NOT NULL COMMENT 'created time',
    updated_at timestamp(6) default CURRENT_TIMESTAMP NOT NULL on update CURRENT_TIMESTAMP COMMENT 'updated time',
    constraint users_pk primary key (id)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_general_ci
    comment = 'tickets list';