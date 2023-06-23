      COMPLEX   ||--o{ UNIT          : has
      UNIT   ||--o{ UNITRESIDENT          : has
      RESIDENT   ||--o{ UNITRESIDENT          : has
      UNIT   ||--o{ LANDLORDOWNUNIT          : has
      LANDLORD   ||--o{ LANDLORDOWNUNIT          : has
      PROPERTYMANAGER   ||--o{ PROPERTYMANAGERUNIT         : has
      UNIT   ||--o{ PROPERTYMANAGERUNIT          : has
      LANDLORD   ||--o{ PROPERTYMANAGERUNIT          : has
      COMPLEX   ||--o{ EXECUTIVECOMMITTEE          : has
      LANDLORD   ||--o{ EXECUTIVECOMMITTEE          : has
      STRATAMANAGER   ||--o{ ASSIGNEDMANAGER          : has
      COMPLEX   ||--o{ ASSIGNEDMANAGER          : has
      CommunityTitleScheme   ||--o{ COMPLEX          : has
      STRATAMANAGER   ||--o{ StrataManagesCTS          : has
      CommunityTitleScheme   ||--o{ StrataManagesCTS          : has
      COMPLEX   ||--o{ AllowedContractors          : has
      CONTRACTOR  ||--o{ AllowedContractors          : has
      PROPERTMANAGER  ||--o{ WorkOrder          : raise
      CONTRACTOR  ||--o{ WorkOrder          : workson
      UNIT  ||--o{ WorkOrder          : worktobeperformed
      RESIDENT  ||--o{ LogUnitIssue           : Called  
      PROPERTYMANAGER  ||--o{ LogUnitIssue           : Called
      ContactPropertyManager ||--o{ WorkOrder          : Called


// there is a sequence when deleting tables
DROP TABLE IF EXISTS unit;
DROP TABLE IF EXISTS complex;
DROP TABLE IF EXISTS resident;

CREATE TABLE IF NOT EXISTS complex (
	complex_id serial PRIMARY KEY,
	name VARCHAR ( 255 ) UNIQUE NOT NULL,
	streetnumber VARCHAR ( 255 ) UNIQUE NOT NULL,
	streetname VARCHAR ( 255 ) UNIQUE NOT NULL,
	upnumber VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS unit (
	unit_id serial PRIMARY KEY,
	fk_complex_id INT NOT NULL,
	name VARCHAR ( 255 ) UNIQUE NOT NULL,
	streetnumber VARCHAR ( 255 ) UNIQUE NOT NULL,
	streetname VARCHAR ( 255 ) UNIQUE NOT NULL,
	upnumber VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
  FOREIGN KEY (fk_complex_id)
        REFERENCES complex (complex_id)
);

CREATE TABLE IF NOT EXISTS resident (
	resident_id serial PRIMARY KEY,
	firstname VARCHAR ( 255 ) UNIQUE NOT NULL,
	surname VARCHAR ( 255 ) UNIQUE NOT NULL,
	dateofbirth VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL
);

