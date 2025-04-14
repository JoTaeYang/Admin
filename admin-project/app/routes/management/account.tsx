import { FormLayout } from "~/components/Form/FormContext";
import { FormInput } from "~/components/Form/FormInput";
import { SaveBtn } from "~/components/Form/SaveBtn";

export default function AccountPage() {
    return (
      <FormLayout>
        <div className="w-[400px] mt-10 ml-8 flex flex-col gap-4">
          <FormInput name="email" placeholder="email" />
          <FormInput name="username" placeholder="username" />
          <FormInput name="role" placeholder="role" />
          <SaveBtn endpoint="/api/save" />
        </div>
      </FormLayout>
    );
  }
  