export const Footer = () => {
  return (
    <footer className="flex justify-between items-center py-4 container mx-auto">
      <span>© {new Date().getFullYear()} Benzen-Games</span>
      <div className="flex space-x-4">
        <a className="hover:underline" href="/terms">
          利用規約
        </a>
        <a className="hover:underline" href="/privacy">
          プライバシーポリシー
        </a>
        <a className="hover:underline" href="/contact">
          お問い合わせ
        </a>
      </div>
    </footer>
  )
}
