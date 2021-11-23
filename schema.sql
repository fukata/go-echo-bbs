create table thread_messages (
    id bigserial primary key,
    message text not null,
    created_at timestamp not null
);