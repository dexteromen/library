import React, { useState, useEffect } from "react";
import "./Home.css";
import Navbar from "../../Components/Navbar/Navbar";
import Card from "../../Components/Card/Card";
import { IoSearch } from "react-icons/io5";
import { Link } from "react-router-dom";
import { getBooks } from "../../API/API";

const books = [
	{
		isbn: "978-3-16-148410-0",
		title: "The Great Gatsby",
		author: "F. Scott Fitzgerald",
		publisher: "Scribner",
		version: "1st Edition",
	},
	{
		isbn: "978-0-7432-7356-5",
		title: "To Kill a Mockingbird",
		author: "Harper Lee",
		publisher: "J.B. Lippincott & Co.",
		version: "1st Edition",
	},
	{
		isbn: "978-0-452-28423-4",
		title: "1984",
		author: "George Orwell",
		publisher: "Secker & Warburg",
		version: "1st Edition",
	},
	{
		isbn: "978-0-7432-7356-6",
		title: "Pride and Prejudice",
		author: "Jane Austen",
		publisher: "T. Egerton",
		version: "1st Edition",
	},
	{
		isbn: "978-0-7432-7356-7",
		title: "The Catcher in the Rye",
		author: "J.D. Salinger",
		publisher: "Little, Brown and Company",
		version: "1st Edition",
	},
	{
		isbn: "978-0-7432-7356-8",
		title: "The Hobbit",
		author: "J.R.R. Tolkien",
		publisher: "George Allen & Unwin",
		version: "1st Edition",
	},
];

export default function Home() {
	const [allbooks, setAllBooks] = useState([]);
	const [searchTerm, setSearchTerm] = useState("");
	const [filter, setFilter] = useState("title");

	const handleSearchChange = (e) => {
		setSearchTerm(e.target.value);
	};

	const handleFilterChange = (e) => {
		setFilter(e.target.value);
	};

	const filteredBooks = books.filter((book) => {
		if (filter === "all") {
			return (
				book.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
				book.author.toLowerCase().includes(searchTerm.toLowerCase()) ||
				book.publisher.toLowerCase().includes(searchTerm.toLowerCase())
			);
		}
		return book[filter].toLowerCase().includes(searchTerm.toLowerCase());
	});

	// useEffect(() => {
	// 	async function fetchBooks() {
	// 		try {
	// 			const res = await getBooks();
	// 			const bookData = res.data.data;
	// 			// const bookData = res;
	// 			// console.log(bookData);
	// 			setAllBooks(bookData);
	// 		} catch (error) {
	// 			console.log(error);
	// 		}
	// 	}
	// 	fetchBooks();
	// }, []);

	return (
		<>
			<Navbar />
			<div className="home-wrapper">
				<div className="logo-centered">
					<img src="/z-library.png" alt="logo-centered" />
					{/* <img src="/image.png" alt="logo-centered" /> */}
				</div>
				<div className="search-box">
					<IoSearch size={"2em"} />
					<input
						type="text"
						placeholder="Search Books"
						value={searchTerm}
						onChange={handleSearchChange}
					/>
					<div className="dropdown-filter">
						<select value={filter} onChange={handleFilterChange}>
							<option value="all">All</option>
							<option value="authors">Author</option>
							<option value="title">Title</option>
							<option value="publisher">Publisher</option>
						</select>
					</div>
				</div>
				<div className="multi-options">
					<Link to="/home">
						<button>Home</button>
					</Link>
					<Link to="/dashboard">
						<button>Dashboard</button>
					</Link>
					<Link to="/create-book">
						<button>Create Book</button>
					</Link>
					<Link to="/create-library">
						<button>Create Library</button>
					</Link>
					<Link to="/temp">
						<button>Temp</button>
					</Link>
				</div>
				<div className="book-cards">
					{filteredBooks.map((book, index) => (
						<Card
							key={index}
							isbn={book.isbn}
							title={book.title}
							author={book.author}
							publisher={book.publisher}
							version={book.version}
						/>
					))}
				</div>
				{/* <div className="book-cards">
					{allbooks.map((book, index) => (
						<Card
							key={index}
							isbn={book.isbn}
							title={book.title}
							author={book.authors}
							publisher={book.publisher}
							version={book.version}
						/>
					))}
				</div> */}
			</div>
		</>
	);
}
