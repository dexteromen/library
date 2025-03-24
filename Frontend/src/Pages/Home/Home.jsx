import React, { useState, useEffect } from "react";
import "./Home.css";
import Navbar from "../../Components/Navbar/Navbar";
import Card from "../../Components/Card/Card";
import { IoSearch } from "react-icons/io5";
import { Link, useNavigate } from "react-router-dom";
import { getBooks, searchBooks, getProfile } from "../../API/API";

export default function Home() {
	const navigate = useNavigate();
	const [allbooks, setAllBooks] = useState([]);
	const [searchTerm, setSearchTerm] = useState("");
	const [filter, setFilter] = useState("title");
	const [user, setUser] = useState({
		name: "",
		email: "",
		role: "",
		contact_number: "",
		libid: 0,
	});
	const [update, setUpdate] = useState(false); // New state to trigger re-render

	const handleSearchChange = (e) => {
		setSearchTerm(e.target.value);
	};

	const handleFilterChange = (e) => {
		setFilter(e.target.value);
	};

	const filteredBooks = allbooks.filter((book) => {
		if (filter === "all") {
			return (
				book.title.toLowerCase().includes(searchTerm.toLowerCase()) ||
				book.authors.toLowerCase().includes(searchTerm.toLowerCase()) ||
				book.publisher.toLowerCase().includes(searchTerm.toLowerCase())
			);
		}
		return book[filter].toLowerCase().includes(searchTerm.toLowerCase());
	});

	useEffect(() => {
		async function fetchBooks() {
			try {
				const res = await getBooks();
				const bookData = res.data.data;
				setAllBooks(bookData);
			} catch (error) {
				console.log(error);
			}
		}
		async function fetchSearchResults() {
			if (searchTerm) {
				try {
					const res = await searchBooks(searchTerm);
					const searchResults = res.data.data;
					setAllBooks(searchResults);
				} catch (error) {
					console.log(error);
				}
			} else {
				fetchBooks();
			}
		}
		fetchSearchResults();

		async function fetchUserData() {
			try {
				const res = await getProfile();
				const userData = res.data.data;
				// console.log(userData);
				setUser({
					name: userData.name,
					email: userData.email,
					role: userData.role,
					contact_number: userData.contact_number,
					libid: userData.lib_id,
				});
			} catch (error) {
				console.error("Token not found in localStorage");
			}
		}
		fetchUserData();
	}, [searchTerm, filter, update]);

	// console.log(allbooks);
	const setUpdates = () => {
		setUpdate(!update);
	};

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
					{user.role === "admin" && (
						<Link to="/dashboard">
							<button>Dashboard</button>
						</Link>
					)}
					{user.role === "owner" && (
						<Link to="/create-book">
							<button>Create Book</button>
						</Link>
					)}
					{user.role === "owner" ||
						(user.role === "reader" && (
							<Link to="/create-library">
								<button>Create Library</button>
							</Link>
						))}
					{user.role === "reader" && (
						<Link to="/manage-books-reader">
							<button>Manage Books</button>
						</Link>
					)}
				</div>
				<div className="book-cards">
					{filteredBooks.length > 0 ? (
						filteredBooks.map((book, index) => (
							<Card
								key={index}
								isbn={book.isbn}
								title={book.title}
								author={book.authors}
								publisher={book.publisher}
								version={book.version}
								lib_id={book.lib_id}
								total_copies={book.total_copies}
								available_copies={book.available_copies}
								user={user}
								updates={setUpdates}
							/>
						))
					) : (
						<p>No books found</p>
					)}
				</div>
			</div>
		</>
	);
}
