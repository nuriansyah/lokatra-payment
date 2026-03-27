CREATE TABLE IF NOT EXISTS event_type (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS events (
    id              UUID PRIMARY KEY,
    name            VARCHAR(255) NOT NULL,
    slug            VARCHAR(255) NOT NULL UNIQUE,
    event_type_id   UUID NOT NULL REFERENCES event_type(id),
    description     TEXT,
    location        VARCHAR(500),
    location_coords JSONB,                                  
    start_date      TIMESTAMPTZ NOT NULL,
    end_date        TIMESTAMPTZ NOT NULL,
    status          VARCHAR(50) NOT NULL DEFAULT 'DRAFT',   
    banner_url      TEXT,
    organizer_name  VARCHAR(255),
    merchant_id     UUID,                                   
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);


CREATE TABLE IF NOT EXISTS event_pricing_policies (
    id              UUID PRIMARY KEY,
    event_id        UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    currency        VARCHAR(3) NOT NULL DEFAULT 'IDR',
    tax_inclusive   BOOLEAN NOT NULL DEFAULT TRUE,
    tax_rate        DECIMAL(5,4) DEFAULT 0,
    pricing_mode    VARCHAR(50) NOT NULL DEFAULT 'MERCHANT_DEFINED',
    notes           TEXT,
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS event_pricing_strategies (
    id              UUID PRIMARY KEY,
    event_id        UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    strategy_type   VARCHAR(50) NOT NULL,
    is_active       BOOLEAN NOT NULL DEFAULT TRUE,
    priority        INT NOT NULL DEFAULT 1,
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS event_pricing_rules (
    id                  UUID PRIMARY KEY,
    strategy_id         UUID NOT NULL REFERENCES event_pricing_strategies(id) ON DELETE CASCADE,
    rule_type           VARCHAR(50) NOT NULL,
    comparator          VARCHAR(10) NOT NULL,
    threshold_value     DECIMAL(15,2) NOT NULL,
    multiplier          DECIMAL(10,4), 
    fixed_adjustment    DECIMAL(15,2), 
    is_active           BOOLEAN NOT NULL DEFAULT TRUE,
    meta_created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by     UUID NOT NULL,
    meta_updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID NOT NULL,
    meta_deleted_at     TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS event_pricing_rule_conditions (
    id              UUID PRIMARY KEY,
    rule_id         UUID NOT NULL REFERENCES event_pricing_rules(id) ON DELETE CASCADE,
    field           VARCHAR(100) NOT NULL,
    operator        VARCHAR(10) NOT NULL,
    value           DECIMAL(15,2) NOT NULL,
    meta_created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by     UUID NOT NULL,
    meta_updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID NOT NULL,
    meta_deleted_at     TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS event_pricing_constraints (
    id                  UUID PRIMARY KEY,
    event_id            UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    constraint_type     VARCHAR(50) NOT NULL,
    value               DECIMAL(15,4) NOT NULL,
    is_active           BOOLEAN NOT NULL DEFAULT TRUE,
    meta_created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by     UUID NOT NULL,
    meta_updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by     UUID NOT NULL,
    meta_deleted_at     TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS venue_spots (
    id              UUID PRIMARY KEY,
    event_id        UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    merchant_id     UUID NOT NULL,                        
    name            VARCHAR(255) NOT NULL,
    spot_type       VARCHAR(50) NOT NULL DEFAULT 'GENERAL', 
    description     TEXT,
    capacity        INT NOT NULL CHECK (capacity > 0),
    sold_count      INT NOT NULL DEFAULT 0 CHECK (sold_count >= 0),
    price           DECIMAL(15,2) NOT NULL CHECK (price >= 0),
    currency        VARCHAR(3) NOT NULL DEFAULT 'IDR',
    status          VARCHAR(50) NOT NULL DEFAULT 'PENDING_APPROVAL', 
    rejection_reason TEXT,
    approved_by     UUID,                                   
    approved_at     TIMESTAMPTZ,
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS commission_schemes (
    id UUID PRIMARY KEY,
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_tiered BOOLEAN NOT NULL DEFAULT FALSE,
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS commission_scheme_participants (
    id UUID PRIMARY KEY,
    scheme_id UUID NOT NULL REFERENCES commission_schemes(id) ON DELETE CASCADE,
    merchant_id UUID NOT NULL, 
    pct DECIMAL(5,4) NOT NULL,
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS commission_rules (
    id UUID PRIMARY KEY,
    scheme_id UUID NOT NULL REFERENCES commission_schemes(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    priority INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS commission_rule_conditions (
    id UUID PRIMARY KEY,
    rule_id UUID NOT NULL REFERENCES commission_rules(id) ON DELETE CASCADE,
    condition_type VARCHAR(50) NOT NULL, 
    operator VARCHAR(20) NOT NULL, 
    value_string TEXT,
    value_number DECIMAL(15,2),
    value_boolean BOOLEAN,
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS commission_rule_actions (
    id UUID PRIMARY KEY,
    rule_id UUID NOT NULL REFERENCES commission_rules(id) ON DELETE CASCADE,
    action_type VARCHAR(50) NOT NULL, 
    value TEXT NOT NULL,
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);

CREATE TYPE status_ticket_enum AS ENUM ('AVAILABLE', 'HELD', 'SOLD', 'CHECKED_IN', 'REFUNDED', 'CANCELLED');
CREATE TABLE IF NOT EXISTS tickets (
    id              UUID PRIMARY,
    ticket_code     VARCHAR(20) NOT NULL UNIQUE, 
    event_id        UUID NOT NULL REFERENCES events(id),
    venue_spot_id   UUID NOT NULL REFERENCES venue_spots(id),
    qr_code_data    TEXT NOT NULL UNIQUE,                
    status          status_ticket_enum NOT NULL DEFAULT 'AVAILABLE',
    price           DECIMAL(15,2) NOT NULL,
    currency        VARCHAR(3) NOT NULL DEFAULT 'IDR',
    purchaser_name  VARCHAR(255),
    purchaser_email VARCHAR(255),
    purchaser_phone VARCHAR(50),
    purchased_at    TIMESTAMPTZ,
    checked_in_at   TIMESTAMPTZ,
    checked_in_by   UUID,                               
    refunded_at     TIMESTAMPTZ,
    refund_reason   TEXT,
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);


CREATE TYPE approval_status_enum AS ENUM ('PENDING', 'APPROVED', 'REJECTED');
CREATE TABLE IF NOT EXISTS approval_requests (
    id              UUID PRIMARY KEY,
    entity_type     VARCHAR(50) NOT NULL,           
    reference_id    UUID NOT NULL,                  
    event_id        UUID NOT NULL REFERENCES events(id),
    merchant_id     UUID NOT NULL,
    status          approval_status_enum NOT NULL DEFAULT 'PENDING',
    submitted_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    reviewed_by     UUID,
    reviewed_at     TIMESTAMPTZ,
    review_notes    TEXT,
    meta_created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_created_by UUID NOT NULL,
    meta_updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    meta_updated_by UUID NOT NULL,
    meta_deleted_at TIMESTAMPTZ
);


-- ============================================================================
-- Updated_at trigger function
-- ============================================================================
CREATE OR REPLACE FUNCTION updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply auto-update triggers
CREATE TRIGGER update_events_updated_at BEFORE UPDATE ON events FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_event_pricing_policies_updated_at BEFORE UPDATE ON event_pricing_policies FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_venue_spots_updated_at BEFORE UPDATE ON venue_spots FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_commission_schemes_updated_at BEFORE UPDATE ON commission_schemes FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_tickets_updated_at BEFORE UPDATE ON tickets FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_approval_requests_updated_at BEFORE UPDATE ON approval_requests FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();


