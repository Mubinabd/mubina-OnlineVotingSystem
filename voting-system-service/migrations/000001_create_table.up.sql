CREATE TABLE if not EXISTS election(
    id uuid PRIMARY key DEFAULT gen_random_uuid(),
    name VARCHAR (60),
    date TIMESTAMP DEFAULT '2024-06-05',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at bigint default 0
);

CREATE TABLE if not EXISTS candidate(
    id uuid PRIMARY key DEFAULT gen_random_uuid(),
    election_id uuid REFERENCES election(id),
    public_id uuid DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at bigint default 0
);


CREATE TABLE if not EXISTS public_vote(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    election_id uuid REFERENCES election(id),
    public_id uuid DEFAULT gen_random_uuid()
);

CREATE TABLE if not EXISTS votes(
    id uuid PRIMARY key DEFAULT gen_random_uuid(),
    candidate_id uuid REFERENCES candidate(id)
);

CREATE TABLE if not EXISTS public(
    id uuid PRIMARY key DEFAULT gen_random_uuid(),
    first_name VARCHAR(60),
    last_name VARCHAR(60),
    birthday TIMESTAMP,
    gender VARCHAR(10),
    nation VARCHAR(100),
    party_id uuid REFERENCES party(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at bigint default 0
);

CREATE TABLE if not EXISTS party(
    id uuid PRIMARY key DEFAULT gen_random_uuid(),
    name VARCHAR(60),
    slogen VARCHAR(255),
    opened_date TIMESTAMP,
    description text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at bigint default 0
);
