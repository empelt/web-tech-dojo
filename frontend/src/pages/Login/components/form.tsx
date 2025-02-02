import { useEffect } from 'react'

import { zodResolver } from '@hookform/resolvers/zod'
import {
  useSignInWithEmailAndPassword,
  useSignInWithGoogle,
} from 'react-firebase-hooks/auth'
import { useForm } from 'react-hook-form'
import { Link } from 'react-router'
import { z } from 'zod'

import LoadingButton from '@/components/LoadingButton'
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
import { toast } from '@/hooks/use-toast'
import { auth } from '@/libs/firebase'

const formSchema = z.object({
  email: z
    .string()
    .nonempty({
      message: 'Email is required.',
    })
    .email({ message: 'Invalid email address.' }),
  password: z.string().nonempty({ message: 'Content is required.' }),
})

const LoginForm = () => {
  const [signInWithGoogle, , googleLoading, googleErr] =
    useSignInWithGoogle(auth)
  const [signInWithEmailAndPassword, , emailLoading, emailErr] =
    useSignInWithEmailAndPassword(auth)

  useEffect(() => {
    if (emailErr || googleErr) {
      toast({
        variant: 'destructive',
        title: 'Uh oh! Something went wrong.',
        description: 'There was a problem with your request.',
      })
    }
  }, [emailErr, googleErr])

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: '',
      password: '',
    },
  })

  function onSubmit(values: z.infer<typeof formSchema>) {
    signInWithEmailAndPassword(values.email, values.password)
  }

  return (
    <Form {...form}>
      <form className="space-y-8" onSubmit={form.handleSubmit(onSubmit)}>
        <div className="flex flex-col gap-6">
          <div className="flex flex-col gap-6">
            <div className="flex flex-col items-center gap-2">
              <a
                className="flex flex-col items-center gap-2 font-medium"
                href="#">
                <img alt="logo" className="w-14" src="/icon.svg" />
                <span className="sr-only">WebTech Dojo</span>
              </a>
              <h1 className="text-xl font-bold">Login</h1>
              <div className="text-center text-sm text-balance text-muted-foreground [&_a]:underline [&_a]:underline-offset-4 hover:[&_a]:text-primary">
                Don&apos;t have an account? <Link to="/signup">Sign up</Link>
              </div>
            </div>
            <div className="flex flex-col gap-6">
              <div className="grid gap-2">
                <FormField
                  control={form.control}
                  name="email"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Email</FormLabel>
                      <FormControl>
                        <Input placeholder="webtech@example.com" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <div className="grid gap-2">
                <FormField
                  control={form.control}
                  name="password"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Password</FormLabel>
                      <FormControl>
                        <Input {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              {emailLoading ? (
                <LoadingButton className="w-full" />
              ) : (
                <Button className="w-full" type="submit">
                  Log in
                </Button>
              )}
            </div>
            <div className="relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t after:border-border">
              <span className="relative z-10 bg-background px-2 text-muted-foreground">
                Or
              </span>
            </div>
          </div>
          {googleLoading ? (
            <LoadingButton className="w-full" />
          ) : (
            <Button
              className="w-full"
              onClick={() => signInWithGoogle()}
              variant="outline">
              <svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path
                  d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z"
                  fill="currentColor"
                />
              </svg>
              Continue with Google
            </Button>
          )}
          <div className="text-balance text-center text-xs text-muted-foreground [&_a]:underline [&_a]:underline-offset-4 hover:[&_a]:text-primary  ">
            By clicking continue, you agree to our{' '}
            <Link to="/terms">Terms of Service</Link> and{' '}
            <Link to="/privacy">Privacy Policy</Link>.
          </div>
        </div>
      </form>
    </Form>
  )
}

export default LoginForm
