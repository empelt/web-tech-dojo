import { zodResolver } from '@hookform/resolvers/zod'
import { useForm } from 'react-hook-form'
import { z } from 'zod'

import { Button } from '@/components/ui/button'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { useToast } from '@/hooks/use-toast'

const formSchema = z.object({
  name: z.string().nonempty({ message: 'Name is required.' }),
  email: z
    .string()
    .nonempty({ message: 'Email is required.' })
    .email({ message: 'Invalid email address.' }),
  content: z.string().nonempty({ message: 'Content is required.' }),
})

const FormKeys = {
  url: import.meta.env.VITE_FORM_URL,
  name: import.meta.env.VITE_FORM_NAME,
  email: import.meta.env.VITE_FORM_EMAIL,
  content: import.meta.env.VITE_FORM_CONTENT,
}

const ContactForm = () => {
  const { toast } = useToast()
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: '',
      email: '',
      content: '',
    },
  })

  function onSubmit(values: z.infer<typeof formSchema>) {
    if (
      !FormKeys.url ||
      !FormKeys.name ||
      !FormKeys.email ||
      !FormKeys.content
    ) {
      console.error('Form keys are not set.')
      return
    }
    const name = encodeURIComponent(values.name)
    const email = encodeURIComponent(values.email)
    const content = encodeURIComponent(values.content)

    fetch(FormKeys.url, {
      method: 'POST',
      mode: 'no-cors',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: `${FormKeys.name}=${name}&${FormKeys.email}=${email}&${FormKeys.content}=${content}`,
    }).then(() => {
      // MEMO: no-cors なので、常に成功する。いつかちゃんとした方法で実装する
      form.reset()
      toast({
        className: 'border-black',
        description: 'お問い合わせを受け付けました。',
      })
    })
  }

  return (
    <Form {...form}>
      <form className="space-y-8" onSubmit={form.handleSubmit(onSubmit)}>
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Name</FormLabel>
              <FormControl>
                <Input placeholder="name" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input placeholder="email" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="content"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Content</FormLabel>
              <FormControl>
                <Textarea placeholder="content" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Submit</Button>
      </form>
    </Form>
  )
}

export default ContactForm
