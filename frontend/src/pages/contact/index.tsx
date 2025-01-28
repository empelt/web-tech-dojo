import ContactForm from './form'

const ContactPage = () => {
  return (
    <div className="container mx-auto p-6 max-w-3xl bg-white rounded-lg shadow-lg mt-6">
      <h1 className="text-2xl font-bold text-center mb-6">お問い合わせ</h1>
      <p className="mb-4">
        以下のフォームから、お問い合わせ内容をご記入ください。
      </p>
      <ContactForm />
    </div>
  )
}

export default ContactPage
