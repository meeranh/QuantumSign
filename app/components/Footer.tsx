import { Facebook, Twitter, Instagram, Linkedin } from "lucide-react"

export default function Footer() {
  return (
    <footer className="bg-gray-100 py-6">
      <div className="container mx-auto">
        <div className="flex justify-center space-x-6">
          <a href="https://www.facebook.com/profile.php?id=61567422373193" target="_blank" rel="noopener noreferrer" className="text-gray-600 hover:text-blue-600">
            <Facebook size={24} />
            <span className="sr-only">Facebook</span>
          </a>
          <a href="https://x.com/QsignAuth" target="_blank" rel="noopener noreferrer" className="text-gray-600 hover:text-blue-400">
            <Twitter size={24} />
            <span className="sr-only">Twitter</span>
          </a>
          <a href="https://www.instagram.com/qsign_99/" target="_blank" rel="noopener noreferrer" className="text-gray-600 hover:text-pink-600">
            <Instagram size={24} />
            <span className="sr-only">Instagram</span>
          </a>
          <a href="https://www.linkedin.com/company/qsign.io" target="_blank" rel="noopener noreferrer" className="text-gray-600 hover:text-blue-700">
            <Linkedin size={24} />
            <span className="sr-only">LinkedIn</span>
          </a>
        </div>
      </div>
    </footer>
  )
}
