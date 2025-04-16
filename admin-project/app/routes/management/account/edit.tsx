import { useEffect, useRef } from "react";
import { useNavigate, useLocation } from "react-router";
import { FormLayout, FormInput, FormDropDown, useFormContext, FormTextField } from "~/components/Form";
import { SaveBtn } from "~/components/Form/SaveBtn";
import BackButton from "~/components/common/BackButton";
import { type Manager } from "~/types/Manager";

function AccountFormContent() {
    const location = useLocation();
    const { setValue } = useFormContext();

    const record = location.state as Manager;
    const initializedRef = useRef(false);
    useEffect(() => {
        if (!initializedRef.current && record) {
            setValue("id", record.id);
            setValue("grade", record.grade);
            setValue("name", record.name);
            initializedRef.current = true;
        }
    }, [record, setValue]);

    return (
        <>
            <FormTextField name="id" placeholder="id" />
            <FormTextField name="name" placeholder="user name" />            
            <FormInput name="grade" placeholder="manager_grade" />            
            <FormDropDown tab= "manager.grade.type" name="grade"/>      
            <SaveBtn endpoint="/api/account/edit/1" />
        </>
    );
}

export default function AccountEditPage() {
    return (
        <div>
            <div className="w-[400px] ml-4 mt-8">
                <div className="flex items-center justify-between mb-6">
                    <h2 className="text-xl font-bold">계정 수정</h2>
                    <BackButton label = "<-"/>
                </div>

                <FormLayout>
                    <AccountFormContent />
                </FormLayout>
            </div>
        </div>

    );
}