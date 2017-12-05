package storage

const SCHEMA_0 = `
	CREATE TABLE InstantDemand(
		Timestamp INTEGER NOT NULL,

		-- in watts
		Demand INTEGER NOT NULL
	);


	-- Summations stores the delivered, received measurements
	-- as sent by the meter
	CREATE TABLE Summations(
		Timestamp INTEGER NOT NULL,

		-- in KWh
		Delivered INTEGER NOT NULL,
		Received INTEGER NOT NULL
	);

	-- increment database version
	PRAGMA user_version=1;
`
