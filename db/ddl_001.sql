create table users(
    id bigserial not null primary key,
    created_at timestamptz not null,
    created_by bigint, -- Only not-null created_by across all db
    updated_at timestamptz,
    updated_by bigint,
    name varchar(100) not null,
    email varchar(100) not null unique,
    hashed_password character varying(60) not null  
);

create table entry_types(
    id bigserial not null primary key,
    created_at timestamptz not null,
    created_by bigint not null,
    updated_at timestamptz,
    updated_by bigint,
    name varchar(20) not null
);

create table user_entry_types(
    id bigserial not null primary key,
    created_at timestamptz not null,
    created_by bigint not null,
    updated_at timestamptz,
    updated_by bigint,
    user_id bigint not null references users(id),
    name varchar(20) not null,
    is_editable boolean not null
);

create table meals(
    id bigserial not null primary key,
    created_at timestamptz not null,
    created_by bigint not null,
    updated_at timestamptz,
    updated_by bigint,
    user_id bigint not null references users(id),
    user_entry_type_id bigint not null references user_entry_types(id),
    date timestamptz
);

create table foods(
    id bigserial not null primary key,
    created_at timestamptz not null,
    created_by bigint not null,
    updated_at timestamptz,
    updated_by bigint,
    meal_id bigint not null references meals(id),
    kcal smallint,
    carbs smallint,
    proteins smallint,
    fats smallint
);