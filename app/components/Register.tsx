'use client';

import { useState, useRef } from 'react';
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Upload, Loader2 } from 'lucide-react';
import Link from 'next/link'

export default function Register() {
	const [username, setUsername] = useState('');
	const [email, setEmail] = useState('');
	const [publicKey, setPublicKey] = useState('');
	const [file, setFile] = useState<File | null>(null);
	const [errorMessage, setErrorMessage] = useState('');
	const [successMessage, setSuccessMessage] = useState('');
	const [isLoading, setIsLoading] = useState(false);
	const [isFileLoading, setIsFileLoading] = useState(false); // Added isFileLoading state
	const fileInputRef = useRef<HTMLInputElement>(null);

	const API_HOST = process.env.NEXT_PUBLIC_API_HOST || 'http://localhost:8000';

	const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
		e.preventDefault();
		setIsLoading(true);

		let kyberPublicKey = publicKey;

		if (file) {
			try {
				const fileContent = await file.text();
				kyberPublicKey = fileContent;
			} catch (err) {
				setErrorMessage('Failed to read the public key file.');
				console.log(err)
				setIsLoading(false);
				return;
			}
		}

		const requestBody = {
			kyberPublicKey,
			username,
			email
		};

		try {
			const response = await fetch(`${API_HOST}/register`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(requestBody)
			});

			if (response.ok) {
				const data = await response.json();
				setSuccessMessage(data.message || 'Registration successful!');
				setErrorMessage('');
			} else {
				const errorData = await response.json();
				if (errorData && errorData.error) {
					setErrorMessage(errorData.error);
				} else if (errorData && errorData.message) {
					setErrorMessage(errorData.message);
				} else {
					setErrorMessage('Internal Server Error');
				}
				setSuccessMessage('');
			}
		} catch (error) {
			console.error('Error:', error);
			setErrorMessage('Internal Server Error');
			setSuccessMessage('');
		} finally {
			setIsLoading(false);
		}
	};

	const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		if (e.target.files && e.target.files.length > 0) {
			const selectedFile = e.target.files[0];
			setFile(selectedFile);
			setIsFileLoading(true); // Set isFileLoading to true

			const reader = new FileReader();
			reader.onload = (event) => {
				const content = event.target?.result as string;
				setPublicKey(content.trim());
				setIsFileLoading(false); // Set isFileLoading to false after reading
			};
			reader.readAsText(selectedFile);
		}
	};

	const handleFileButtonClick = () => {
		fileInputRef.current?.click();
	};

	return (
		<div className="flex justify-center items-center min-h-screen bg-gradient-to-r from-gray-100 to-gray-200">
			<div className="bg-white p-10 rounded-3xl shadow-2xl w-full max-w-lg">
				<h2 className="text-3xl font-bold text-center mb-6 text-gray-800">
					SIGN UP
				</h2>

				{successMessage && (
					<div className="mb-4 text-green-500 text-center font-semibold">
						{successMessage}
					</div>
				)}

				{errorMessage && (
					<div className="mb-4 text-red-500 text-center font-semibold">
						{errorMessage}
					</div>
				)}

				<form onSubmit={handleSubmit}>
					<div className="space-y-6">
						<div>
							<Label
								htmlFor="username"
								className="block text-gray-700 font-semibold mb-1"
							>
								Enter Your Username
							</Label>
							<Input
								id="username"
								type="text"
								placeholder="Username"
								value={username}
								onChange={(e) => setUsername(e.target.value)}
								className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-400"
							/>
						</div>

						<div>
							<Label
								htmlFor="email"
								className="block text-gray-700 font-semibold mb-1"
							>
								Enter Your Email
							</Label>
							<Input
								id="email"
								type="email"
								placeholder="Email"
								value={email}
								onChange={(e) => setEmail(e.target.value)}
								className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-400"
							/>
						</div>

						<div>
							<Label
								htmlFor="publicKey"
								className="block text-gray-700 font-semibold mb-1"
							>
								Enter Your Public Key
							</Label>
							<ul className="mt-2 ml-2 list-disc list-inside text-gray-600">
									<li className="text-sm text-gray-500 mb-2">
										<Link href='/gen-keypair' className='text-blue-600 underline italic'>Click me</Link> for a public key!
									</li>
								</ul>
							<Input
								id="publicKey"
								type="text"
								placeholder="Public Key"
								value={publicKey}
								onChange={(e) => setPublicKey(e.target.value)}
								className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-400"
							/>
						</div>
					</div>

					<div className="text-center my-8">
						<p className="text-gray-500 mb-4 text-sm">OR</p>
						<Label className="block font-bold text-gray-700 mb-2">
							UPLOAD YOUR PUBLIC KEY FILE
						</Label>
						<div className="bg-gray-50 rounded-lg p-6 shadow-md flex justify-center">
							<input
								ref={fileInputRef}
								type="file"
								accept=".pub"
								onChange={handleFileChange}
								className="hidden"
							/>
							<Button
								type="button"
								onClick={handleFileButtonClick}
								className="bg-blue-500 text-white px-6 py-3 rounded-full hover:bg-blue-600 transition-all duration-300 flex items-center justify-center"
								disabled={isFileLoading} // Use isFileLoading for button disable state
							>
								{isFileLoading ? (
									<>
										<Loader2 className="mr-2 h-5 w-5 animate-spin" />
										Reading file...
									</>
								) : (
									<>
										<Upload className="mr-2 h-5 w-5" />
										{file ? file.name : 'Upload File'}
									</>
								)}
							</Button>
						</div>
					</div>

					<div className="mt-6 text-center">
						<Button
							type="submit"
							className="bg-black text-white w-full py-3 rounded-full text-lg font-bold hover:bg-gray-900 focus:ring-2 focus:ring-yellow-400 flex items-center justify-center"
							disabled={isLoading}
						>
							{isLoading ? (
								<>
									<Loader2 className="mr-2 h-4 w-4 animate-spin" />
									Registering...
								</>
							) : (
								'REGISTER'
							)}
						</Button>
					</div>
				</form>
			</div>
		</div>
	);
}
