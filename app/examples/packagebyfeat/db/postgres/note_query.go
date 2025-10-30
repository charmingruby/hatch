package postgres

const (
	createNote   = "create note"
	findNoteByID = "find note by id"
	listNotes    = "list notes"
	saveNote     = "save note"
)

func noteQueries() map[string]string {
	return map[string]string{
		createNote: `INSERT INTO notes 
			(id, title, content, archived, created_at, updated_at) 
			VALUES ($1, $2, $3, $4, $5, $6)`,
		findNoteByID: `SELECT * FROM notes WHERE id = $1`,
		listNotes:    `SELECT * FROM notes`,
		saveNote: `UPDATE notes 
			SET archived = $1, updated_at = $2  
			WHERE id = $3`,
	}
}
