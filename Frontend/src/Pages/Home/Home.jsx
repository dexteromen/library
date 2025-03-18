import React, { useState } from "react";
import "./Home.css";
import Navbar from "../../Components/Navbar/Navbar";
import Card from "../../Components/Card/Card";
import { IoSearch } from "react-icons/io5";

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

	return (
		<>
			<Navbar />
			<div className="home-wrapper">
				<div className="logo-centered">
					<img src="/z-library.png" alt="logo-centered" />
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
							<option value="author">Author</option>
							<option value="title">Title</option>
							<option value="publisher">Publisher</option>
						</select>
					</div>
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
			</div>
		</>
	);
}
