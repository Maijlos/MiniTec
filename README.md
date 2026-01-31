# Health check system

This project was created for PB173 course at Masaryk University.

It is a simple UI table where u can display state of machine labeled under project and then through CSV file send data which will be stored in DB under choosen project.

Data in CSV needs to be in specific format with specific headers.

(more in demo presentation)

## How to run

To run DB and backend you need to set `.env` to work properly.

Then you can run:
```
docker compose up -d
```

For frontend you need to run and also set `.env`:
```
cd ./frontend
npm run dev
```
