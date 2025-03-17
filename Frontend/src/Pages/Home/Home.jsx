import React from "react";
import "./Home.css";
import Navbar from "../../Components/Navbar/Navbar";
import Card from "../../Components/Card/Card";
import { IoSearch } from "react-icons/io5";

const books = [
	{
		bookImage: "Book Image 1",
		desc: "Description 1",
	},
	{
		bookImage: "Book Image 2",
		desc: "Description 2",
	},
	{
		bookImage: "Book Image 3",
		desc: "Description 3",
	},
	{
		bookImage: "Book Image 4",
		desc: "Description 4",
	},
	{
		bookImage: "Book Image 5",
		desc: "Description 5",
	},
	{
		bookImage: "Book Image 6",
		desc: "Description 6",
	},
];

export default function Home() {
	return (
		<>
			<Navbar />
			<div className="home-wrapper">
				<div className="logo-centered">
					<img src="/z-library.png" alt="logo-centered" />
				</div>
				<div className="search-box">
					<IoSearch size={"2em"} />
					<input type="text" placeholder="Search Books" />
					<div className="dropdown-filter">
						<select>
							<option value="authors">author</option>
							<option value="title" selected>
								Title
							</option>
							<option value="publisher">Publisher</option>
						</select>
					</div>
				</div>

				<div className="book-cards">
					{books.map((book, index) => (
						<Card
							key={index}
							bookImage={book.bookImage}
							desc={book.desc}
						/>
					))}
				</div>
			</div>
		</>
	);
}
