db.createUser(
    {
        user: "DASistemas",
        pwd: "ex-especial567",
        roles: [
            {
                role: "readWrite",
                db: "database_for_persons"
            }
        ]
    }
);