CREATE TABLE party (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50),
    slogan VARCHAR(50),
    open_date TIMESTAMP NOT NULL,
    description TEXT NOT NULL
);


CREATE TABLE public (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    birthday TIMESTAMP NOT NULL,
    gender VARCHAR(1) NOT NULL,
    nation VARCHAR(30) NOT NULL,
    party_id UUID REFERENCES party(id)
);

CREATE TABLE election (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    date TIMESTAMP NOT NULL
);


CREATE TABLE candidate (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    election_id UUID  UNIQUE REFERENCES election(id),
    public_id UUID  UNIQUE,
    party_id UUID
);


CREATE TABLE public_vote (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    election_id UUID  UNIQUE REFERENCES election(id),
    public_id UUID UNIQUE
);

CREATE TABLE vote (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    candidate_id UUID REFERENCES candidate(id)
);