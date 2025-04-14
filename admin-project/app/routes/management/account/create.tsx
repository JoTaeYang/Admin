import { FormLayout } from "~/components/Form/FormContext";
import { FormInput } from "~/components/Form/FormInput";
import { SaveBtn } from "~/components/Form/SaveBtn";
import BackButton from "~/components/common/BackButton";
import { useNavigate } from "react-router";

export default function AccountCreatePage() {
    const navigate = useNavigate();

    return (
        <div className="w-[400px] ml-4 mt-8">
            <div className="flex items-center justify-between mb-6">
                <h2 className="text-xl font-bold">계정 수정</h2>
                <BackButton label="<-" />
            </div>

            <FormLayout>
                <FormInput name="email" placeholder="이메일" />
                <FormInput name="username" placeholder="사용자 이름" />
                <FormInput name="role" placeholder="권한" />
                <SaveBtn
                    endpoint="/api/account/create"
                    onSuccess={() => navigate("..")}
                />
            </FormLayout>
        </div>
    );
}
