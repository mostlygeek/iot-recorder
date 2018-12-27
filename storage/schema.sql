-- instant demand records
CREATE TABLE demand (
    id              INTEGER NOT NULL PRIMARY KEY,
	timestamp       INTEGER NOT NULL,

	-- in watts
	demand          MEDIUMINT NOT NULL
);
CREATE INDEX DEMAND_TS on demand(timestamp);

-- summations stores the delivered, received measurements
-- as sent by the meter
CREATE TABLE summations (
    id              INTEGER NOT NULL PRIMARY KEY,
	timestamp       INTEGER NOT NULL,

	-- in Wh
	delivered       FLOAT NOT NULL,
	received        FLOAT NOT NULL
);

CREATE INDEX SUM_TS on summations(timestamp);

-- increment database version
PRAGMA user_version=1;
