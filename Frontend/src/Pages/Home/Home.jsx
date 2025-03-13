import React from "react";
import "./Home.css";
import Navbar from "../../Components/Navbar/Navbar";
import Card from "../../Components/Card/Card";

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
];

export default function Home() {
	return (
		<>
			<Navbar />
			<div className="book-cards">
				{books.map((book, index) => (
					<Card
						key={index}
						bookImage={book.bookImage}
						desc={book.desc}
					/>
				))}
			</div>
		</>
	);
}
