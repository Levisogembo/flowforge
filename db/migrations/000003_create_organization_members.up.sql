CREATE TYPE org_role AS ENUM ('OWNER', 'ADMIN', 'MEMBER', 'VIEWER');

CREATE TABLE organization_members (
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role org_role NOT NULL DEFAULT 'MEMBER',
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (organization_id, user_id)
);