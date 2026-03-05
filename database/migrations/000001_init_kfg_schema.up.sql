-- Knowledge Flow Graph Nodes
CREATE TABLE kfg_nodes (
    node_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    signal_type VARCHAR(255) NOT NULL,
    current_state JSONB NOT NULL,
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    compliance_status VARCHAR(50) DEFAULT 'PENDING'
);

-- Audit log for Agent decisions (NIST 800-53 requirement)
CREATE TABLE agent_decisions (
    decision_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    node_id UUID REFERENCES kfg_nodes(node_id),
    agent_type VARCHAR(100), -- 'research' or 'compliance'
    rationale TEXT,
    confidence_score FLOAT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_kfg_signal_type ON kfg_nodes(signal_type);
