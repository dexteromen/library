import React from "react";
import "./Card.css";

function Card({ isbn, title, author, publisher, version }) {
    return (
        <>
            <div className="card">
                <table className="book-details">
                    <tbody>
                        <tr>
                            <td className="table-label">ISBN</td>
                            <td>{isbn}</td>
                        </tr>
                        <tr>
                            <td className="table-label">Title</td>
                            <td>{title}</td>
                        </tr>
                        <tr>
                            <td className="table-label">Author</td>
                            <td>{author}</td>
                        </tr>
                        <tr>
                            <td className="table-label">Publisher</td>
                            <td>{publisher}</td>
                        </tr>
                        <tr>
                            <td className="table-label">Version</td>
                            <td>{version}</td>
                        </tr>
                    </tbody>
                </table>
                <button className="btn-request">Request</button>
            </div>
        </>
    );
}

export default Card;