
// Create a record
db.Debug().Create(user)
// Also we can use save that will return primary key
db.Debug().Save(newUser)

// transaction
tx := db.Begin()
err := tx.Create(&user).Error
if err != nil {
 tx.Rollback()
}
tx.Commit()